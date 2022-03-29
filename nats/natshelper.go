package nats

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/pnocera/res-gomodel/config"
	"github.com/pnocera/res-gomodel/types"
)

//Config struct using viper
type NatsHelper struct {
	conf *config.Config
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
	//defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	c.js = js

	return &c, nil
}

func (nh *NatsHelper) AddStream(name string) error {
	jsinfo, _ := nh.js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".>"},
		MaxAge:   32 * time.Minute,
	})

	bytes, _ := json.MarshalIndent(jsinfo, "", " ")
	log.Println("Initialized stream  " + name)
	log.Println(string(bytes))

	return nil
}

func (nh *NatsHelper) Publish(subject string, payload interface{}) (string, error) {

	id := uuid.New().String()
	message := types.Message{
		ID:      id,
		Payload: payload,
	}
	messagejson, _ := json.Marshal(message)

	_, err := nh.js.Publish(subject, messagejson)

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
