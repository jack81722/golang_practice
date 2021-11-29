package prepare

import (
	"pubsubexp/adapter"
	"pubsubexp/service"
)

type Prepare struct {
	pubsub  *adapter.PubSub
	expServ *service.ExpService
}

func (p *Prepare) DoAll() (err error) {
	if err = p.SetPubSub(); err != nil {
		return
	}
	if err = p.SetExpService(); err != nil {
		return
	}
	return
}

func (p *Prepare) Close() (err error) {
	p.pubsub.Close()
	return
}

func (p *Prepare) SetPubSub() (err error) {
	if p.pubsub != nil {
		return
	}
	topic := "exp_topic"
	sub := "exp_sub"
	proj := "lcwp-jack"
	p.pubsub, err = adapter.NewPubSub(topic, sub, proj)
	return
}

func (p *Prepare) PubSub() *adapter.PubSub {
	return p.pubsub
}

func (p *Prepare) SetExpService() (err error) {
	if p.expServ != nil {
		return
	}
	p.expServ = service.NewExpService(p.PubSub())
	return
}

func (p *Prepare) ExpService() *service.ExpService {
	return p.expServ
}
