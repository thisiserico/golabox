package subscriber

import (
	"log"

	"github.com/thisiserico/golabox/domain/event"
	"github.com/thisiserico/golabox/domain/service"
	"github.com/thisiserico/golabox/eventbus"
)

type Subscriber struct {
	bus           chan eventbus.Event
	domainService *service.Service
}

func NewSubscriber(ch chan eventbus.Event, srv *service.Service) *Subscriber {
	return &Subscriber{
		bus:           ch,
		domainService: srv,
	}
}

func (s *Subscriber) HandleEvents() {
	for ev := range s.bus {
		log.Printf(
			"handling event %s:%s triggered at %s\n",
			ev.EventName(),
			ev.EventID(),
			ev.TriggeredAt(),
		)

		switch ev.(type) {
		case *event.OrderWasPaid:
			typedEv := ev.(*event.OrderWasPaid)
			s.domainService.DecreaseProductsStock(typedEv.Order)

		default:
			log.Printf("unknown event type %s\n", ev.EventName())
		}

		// Save the event somewhere!
		log.Printf("event %s:%s handled\n", ev.EventName(), ev.EventID())
	}
}
