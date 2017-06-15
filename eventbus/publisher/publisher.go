package publisher

import (
	"log"

	"github.com/thisiserico/golabox/eventbus"
)

type Publisher struct {
	bus chan eventbus.Event
}

func NewPublisher(ch chan eventbus.Event) *Publisher {
	return &Publisher{
		bus: ch,
	}
}

func (p *Publisher) Publish(ev eventbus.Event) {
	log.Printf("publishing event %s:%s\n", ev.EventName(), ev.EventID())
	p.bus <- ev
}
