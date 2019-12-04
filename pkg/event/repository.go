package event

import (
	"context"

	"github.com/google/uuid"
	"github.com/tmtx/erp/pkg/bus"
	"go.mongodb.org/mongo-driver/bson"
)

type UUID struct {
	uuid.UUID
}

type Event struct {
	bus.Message
	EntityId *UUID `bson:"entity_id"`
}

type Filter bson.M

func New(key bus.MessageKey, params bus.MessageParams) Event {
	return Event{
		bus.Message{
			Key:    key,
			Params: params,
			Type:   bus.EventMessage,
		},
		nil,
	}
}

func NewWithId(key bus.MessageKey, params bus.MessageParams) Event {
	e := New(key, params)
	e.CreateNewId()
	return e
}

type Repository interface {
	Store(ctx context.Context, e Event) error
	FindOneWithFilter(ctx context.Context, filter Filter, result *Event) error
	FindAllWithFilter(ctx context.Context, filter Filter) ([]Event, error)
}

func (e Event) CreateNewId() {
	e.EntityId = &UUID{uuid.New()}
}
