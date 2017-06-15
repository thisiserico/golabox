package eventbus

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Event interface {
	TriggeredAt() time.Time
	EventID() uuid.UUID
	EventName() string
}
