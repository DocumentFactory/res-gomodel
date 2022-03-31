package natsutils

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/pnocera/res-gomodel/config"
)

//Config struct using viper
type NatsHelper struct {
	conf *config.Config
	nc   *nats.Conn
	js   nats.JetStreamContext
	subs []*nats.Subscription
}

type SubscribeInvocationHandler func(ctx context.Context, msg *nats.Msg) error

//New Create a new config
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

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	c.js = js

	c.addStream("feedback", 30*time.Second)
	c.addStream("tasks", 32*time.Minute)

	return &c, nil
}

func (nh *NatsHelper) addStream(name string, maxage time.Duration) error {
	conf := &nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".>"},
		MaxAge:   maxage,
	}
	jsinfo, err := nh.js.AddStream(conf)

	if err != nil {
		jsinfo, err = nh.js.UpdateStream(conf)
	}

	bytes, _ := json.MarshalIndent(jsinfo, "", " ")
	log.Println("Initialized stream  " + name)
	log.Println(string(bytes))

	return nil
}

func (nh *NatsHelper) Broadcast(subject string, payload interface{}) error {

	messagejson, _ := json.Marshal(payload)

	_, err := nh.js.Publish(subject, messagejson)

	return err
}

func (nh *NatsHelper) Close() {
	nh.nc.Close()
}

func (nh *NatsHelper) Publish(subject string, payload interface{}) (string, error) {

	id := uuid.New().String()

	messagejson, _ := json.Marshal(payload)

	_, err := nh.js.Publish(subject, messagejson, nats.MsgId(id), nats.AckWait(1*time.Second))

	if err != nil {
		return "", err
	}

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

		go func() {
			for {
				msgs, _ := sub.Fetch(poolsize, nats.MaxWait(time.Second))
				log.Printf("Got %d messages to process\n", len(msgs))
				for _, msg := range msgs {
					msg.Ack()
					// meta, _ := msg.Metadata()
					// go log.Printf(
					// 	"got msg %d/%d on subject %s",
					// 	meta.Sequence.Stream,
					// 	meta.Sequence.Consumer,
					// 	msg.Subject,
					// )
					err = fn(context.Background(), msg)
					if err != nil {
						log.Printf("Error processing message %v", err)
					}
				}
				log.Printf("Processed %d messages. Going for another round trip...\n", len(msgs))
			}

		}()

	} else {
		sub, err = nh.js.Subscribe(subject, func(msg *nats.Msg) {
			msg.Ack()
			err = fn(context.Background(), msg)
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
