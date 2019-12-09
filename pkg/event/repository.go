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
	Key      bus.MessageKey    `bson:"key"`
	Params   bus.MessageParams `bson:"params"`
	EntityId *UUID             `bson:"entity_id"`
}

type Filter bson.M

func New(key bus.MessageKey, params bus.MessageParams, entityId *UUID) Event {
	e := Event{
		key,
		params,
		entityId,
	}
	e.Params = params
	return e
}

type Repository interface {
	Store(ctx context.Context, e Event) error
	FindOneWithFilter(ctx context.Context, filter Filter) (Event, error)
	FindAllWithFilter(ctx context.Context, filter Filter) ([]Event, error)
}

func CreateNewEntityId() *UUID {
	return &UUID{uuid.New()}
}
