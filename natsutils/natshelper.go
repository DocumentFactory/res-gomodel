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
)

type NatsHelper struct {
	conf *config.Config
	nc   *nats.Conn
	js   nats.JetStreamContext
	subs []*nats.Subscription
}

type SubscribeInvocationHandler func(ctx context.Context, msg *nats.Msg) error

type CancelWorkflowHandler func(ctx context.Context, msg *nats.Msg) bool

//New Create a new NatsHelper
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
	c.addStream("cancel", 30*time.Second)

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
	nh.nc.Flush()
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

func (nh *NatsHelper) AddSubscribeHandler(pool string, poolsize int, subject string, fn SubscribeInvocationHandler, chk CancelWorkflowHandler) error {
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
						if !chk(ctx, m) {
							err = fn(ctx, m)
							if err != nil {
								log.Printf("Error processing message %v", err)
							}
						}

						wg.Done()
					}(msg)
				}
				wg.Wait()
				log.Printf("Processed %d messages. Going for another round trip...\n", len(msgs))
			}

		}(sub)

	} else {
		sub, err = nh.js.Subscribe(subject, func(msg *nats.Msg) {
			msg.Ack()
			ctx := context.Background()
			if !chk(ctx, msg) {
				err = fn(ctx, msg)
				if err != nil {
					log.Printf("Error processing message %v", err)
				}
			}
		})
		if err != nil {
			return err
		}
	}
	return nil
}
