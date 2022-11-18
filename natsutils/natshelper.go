package natsutils

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/minio/sio"
	"github.com/nats-io/nats.go"
	"github.com/pnocera/res-gomodel/config"
)

type NatsHelper struct {
	conf *config.Config
	nc   *nats.Conn
	js   nats.JetStreamContext
	// subs  []*nats.Subscription
	store nats.ObjectStore
}

type SubscribeInvocationHandler func(ctx context.Context, msg *nats.Msg) error

type CancelWorkflowHandler func(ctx context.Context, msg *nats.Msg) bool

// New Create a new NatsHelper
func NewNatsHelper(conf *config.Config) (*NatsHelper, error) {
	c := NatsHelper{
		conf: conf,
	}
	opt, err := nats.NkeyOptionFromSeed(conf.NatsNKeyPath())
	if err != nil {
		return nil, err
	}
	nc, err := nats.Connect(conf.NatsURL(), opt)

	if err != nil {
		return nil, err
	}

	c.nc = nc

	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	c.js = js

	c.addStream("feedback", 30*time.Second)
	c.addStream("tasks", 32*time.Minute)
	c.addStream("cancel", 30*time.Second)

	c.store, err = js.ObjectStore("tempfiles")
	if err != nil {

		c.store, err = c.js.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket:      "tempfiles",
			Description: "Temporary files",
			TTL:         30 * time.Second,
			MaxBytes:    1024 * 1024 * 1024 * 1024,
			Storage:     nats.FileStorage,
			Replicas:    1,
		})
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (nh *NatsHelper) PutBytes(key string, data []byte) (*nats.ObjectInfo, error) {

	return nh.store.PutBytes(key, data)

}

func (nh *NatsHelper) Put(name string, masterkey string, reader io.Reader) ([32]byte, error) {
	nonce, err := GetNonce()
	if err != nil {
		return [32]byte{}, err
	}

	key, err := GetKey(masterkey, nonce)
	if err != nil {
		return [32]byte{}, err
	}

	encrypted, err := sio.EncryptReader(reader, sio.Config{Key: key[:]})
	if err != nil {
		return [32]byte{}, err
	}

	_, err = nh.store.Put(&nats.ObjectMeta{
		Name: name,
	}, encrypted)

	if err != nil {
		return [32]byte{}, err
	}

	return nonce, nil

}

func (nh *NatsHelper) GetBytes(key string) ([]byte, error) {

	return nh.store.GetBytes(key)

}

func (nh *NatsHelper) Get(masterkey string, nonce [32]byte, name string) (io.Reader, error) {

	key, err := GetKey(masterkey, nonce)
	if err != nil {
		return nil, err
	}

	obj, err := nh.store.Get(name)
	if err != nil {
		return nil, err
	}

	decrypted, err := sio.DecryptReader(obj, sio.Config{Key: key[:]})
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func (nh *NatsHelper) Delete(key string) error {

	return nh.store.Delete(key)

}

func (nh *NatsHelper) addStream(name string, maxage time.Duration) error {
	conf := &nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".>"},
		MaxAge:   maxage,
	}
	jsinfo, err := nh.js.AddStream(conf)

	if err != nil {
		jsinfo, _ = nh.js.UpdateStream(conf)
	}

	bytes, _ := json.MarshalIndent(jsinfo, "", " ")
	log.Println("Initialized stream  " + name)
	log.Println(string(bytes))

	return nil
}

func (nh *NatsHelper) Broadcast(subject string, payload interface{}) error {

	messagejson, _ := json.Marshal(payload)

	_, err := nh.js.PublishAsync(subject, messagejson)
	select {
	case <-nh.js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		log.Println("Did not resolve in time")
	}

	return err
}

func (nh *NatsHelper) Close() {
	nh.nc.Flush()
	nh.nc.Close()
}

func (nh *NatsHelper) Publish(subject string, payload interface{}) (string, error) {

	id := uuid.New().String()

	messagejson, _ := json.Marshal(payload)

	nh.js.PublishAsync(subject, messagejson, nats.MsgId(id))
	select {
	case <-nh.js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		log.Println("Did not resolve in time")
	}
	// _, err := nh.js.Publish(subject, messagejson, nats.MsgId(id), nats.AckWait(1*time.Second))

	// if err != nil {
	// 	return "", err
	// }

	return id, nil
}

func (nh *NatsHelper) AddSubscribeHandler(pool string, poolsize int, subject string, fn SubscribeInvocationHandler) error {
	var sub *nats.Subscription
	var err error
	if poolsize < 1 {
		poolsize = 5
	}
	if pool != "" {

		sub, err = nh.js.PullSubscribe(subject, pool)
		if err != nil {
			return err
		}

		go func(subscription *nats.Subscription) {
			for subscription.IsValid() {

				msgs, _ := subscription.Fetch(poolsize, nats.MaxWait(1*time.Second))

				numDigesters := len(msgs)
				if numDigesters == 0 {
					continue
				}

				var wg sync.WaitGroup
				log.Printf("Got %d messages to process\n", numDigesters)
				wg.Add(numDigesters)
				for _, msg := range msgs {
					msg.Ack()
					go func(m *nats.Msg) {
						ctx := context.Background()
						err = fn(ctx, m)
						if err != nil {
							log.Printf("Error processing message %v", err)
						}

						wg.Done()
					}(msg)
				}
				wg.Wait()
				log.Printf("Processed %d messages. Going for another round trip...\n", len(msgs))
			}

		}(sub)

	} else {
		_, err = nh.js.Subscribe(subject, func(msg *nats.Msg) {
			msg.Ack()
			ctx := context.Background()
			err = fn(ctx, msg)
			if err != nil {
				log.Printf("Error processing message %v", err)
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}
