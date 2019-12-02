package event

import (
	"context"

	"github.com/tmtx/erp/pkg/bus"
	"go.mongodb.org/mongo-driver/bson"
)

type Event struct {
	bus.Message
}

type Filter bson.M

func New(key bus.MessageKey, params bus.MessageParams) Event {
	return Event{
		bus.Message{
			Key:    key,
			Params: params,
			Type:   bus.EventMessage,
		},
	}
}

type Repository interface {
	Store(ctx context.Context, e Event) error
	FindOneWithFilter(ctx context.Context, filter Filter, result *Event) error
}
