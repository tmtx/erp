package app

import (
	"github.com/tmtx/res-sys/pkg/bus"
	"github.com/tmtx/res-sys/pkg/event"
)

type BasicService struct {
	CommandBus      bus.MessageBus
	EventRepository event.Repository
}

func NewBasicService(commandBus bus.MessageBus, eventRepository event.Repository) BasicService {
	return BasicService{
		commandBus,
		eventRepository,
	}
}

func RegisterCommandSubscribers(subscribers []CommandSubscriber) {
	for _, s := range subscribers {
		s.RegisterCommandCallbacks()
	}
}
