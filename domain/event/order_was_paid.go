package event

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/thisiserico/golabox/domain"
)

type OrderWasPaid struct {
	id          uuid.UUID
	triggeredAt time.Time
	Order       *domain.Order
}

func NewOrderWasPaidEvent(o *domain.Order) *OrderWasPaid {
	return &OrderWasPaid{
		id:          uuid.NewV4(),
		triggeredAt: time.Now().UTC(),
		Order:       o,
	}
}

func (ev *OrderWasPaid) TriggeredAt() time.Time {
	return ev.triggeredAt
}

func (ev *OrderWasPaid) EventID() uuid.UUID {
	return ev.id
}

func (ev *OrderWasPaid) EventName() string {
	return "order_was_paid"
}
