package service

import (
	"context"
	"fmt"
	"log"
	"pubsubexp/adapter"

	"cloud.google.com/go/pubsub"
)

type ExpService struct {
	cancel context.CancelFunc
}

func NewExpService(pubsub *adapter.PubSub) *ExpService {
	serv := &ExpService{}
	ctx, cancel := context.WithCancel(context.TODO())
	serv.cancel = cancel
	go func() error {
		e := pubsub.Subscribe(ctx, serv.SayHello)
		return e
	}()
	go func() error {
		p, e := adapter.NewPubSub("exp_topic", "exp_sub2", "lcwp-jack")
		if e != nil {
			return e
		}
		e = p.Subscribe(context.Background(), serv.SayBye)
		return e
	}()
	log.Println("service new")
	return serv
}

func (e *ExpService) Close() {
	e.cancel()
}

func (e *ExpService) SayHello(ctx context.Context, msg *pubsub.Message) {
	fmt.Println("Hello ", string(msg.Data))
	msg.Ack()
}

func (e *ExpService) SayBye(ctx context.Context, msg *pubsub.Message) {
	fmt.Println("Bye ", string(msg.Data))
	msg.Ack()
}
