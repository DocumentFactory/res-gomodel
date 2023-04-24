package natsutils

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/pnocera/res-gomodel/config"
	"github.com/pnocera/res-gomodel/types"
)

type NatsHelper struct {
	conf *config.Config
	nc   *nats.Conn
	js   nats.JetStreamContext
	wfkv nats.KeyValue
}

type SubscribeInvocationHandler func(ctx context.Context, msg *nats.Msg) error

type WatchKVHandler func(ctx context.Context, value nats.KeyValueEntry) error

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
	c.addStream("cleanup", 30*time.Second)

	err = c.addKV("workflows", 120*time.Minute)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (nh *NatsHelper) addKV(name string, maxage time.Duration) error {
	var kv nats.KeyValue
	var err error
	if stream, _ := nh.js.StreamInfo(name); stream == nil {
		// A key-value (KV) bucket is created by specifying a bucket name.
		kv, err = nh.js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket:  name,
			History: 5,
			TTL:     maxage,
		})
	} else {
		kv, err = nh.js.KeyValue(name)
	}
	if err == nil {
		log.Println("Initialized kv store  " + name)
		nh.wfkv = kv
	}
	return err
}

func (nh *NatsHelper) addStream(name string, maxage time.Duration) error {
	conf := &nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".>"},
		MaxAge:   maxage,
	}
	_, err := nh.js.AddStream(conf)

	if err != nil {
		_, _ = nh.js.UpdateStream(conf)
	}

	log.Println("Initialized stream  " + name)

	return nil
}

func (nh *NatsHelper) PutWfKV(runid string, payload types.WFKeyVal) error {
	messagejson, _ := json.Marshal(payload)
	_, err := nh.wfkv.Put(runid, messagejson)

	if err != nil {
		log.Printf("Error put workflow %s, error : %v", runid, err)
	}

	return err
}

func (nh *NatsHelper) GetWfKVAll() ([]types.WFKeyVal, error) {
	var result []types.WFKeyVal = make([]types.WFKeyVal, 0)

	var err error

	keys, err := nh.wfkv.Keys()

	if err != nil {
		log.Printf("Error get workflow keys, error : %v", err)
		return result, err
	}

	for _, key := range keys {
		msg, _ := nh.wfkv.Get(key)
		if msg != nil {
			var keyval types.WFKeyVal
			err = json.Unmarshal(msg.Value(), &keyval)
			if err == nil {
				if keyval.Status > 0 {
					result = append(result, keyval)
				}
			}
		} else {
			log.Printf("Error get workflow %s, error : %v", key, err)
		}
	}

	return result, err
}

func (nh *NatsHelper) GetWfKV(key string) (*types.WFKeyVal, error) {
	msg, err := nh.wfkv.Get(key)
	if err != nil {
		return nil, err
	}
	var keyval types.WFKeyVal
	err = json.Unmarshal(msg.Value(), &keyval)
	if err != nil {
		return nil, err
	}
	return &keyval, nil
}

func (nh *NatsHelper) DeleteWfKV(key string) error {
	return nh.wfkv.Delete(key)
}

func (nh *NatsHelper) WatchWfKV(fn WatchKVHandler) error {
	w, err := nh.wfkv.WatchAll()
	if err != nil {
		return err
	}

	go func() {

		for kve := range w.Updates() {
			if kve != nil {
				fn(context.Background(), kve)
			}
		}
	}()
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

// func (nh *NatsHelper) AddWatchKVHandler(kv string, fn WatchKVHandler) error {

// 	var err error
// 	_, err = nh.js.Subscribe(kv, func(msg *nats.Msg) {
// 		msg.Ack()
// 		ctx := context.Background()
// 		err = fn(ctx, msg.Data)
// 		if err != nil {
// 			log.Printf("Error processing message %v", err)
// 		}
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
