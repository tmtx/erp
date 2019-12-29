package bus

import (
	"fmt"

	"github.com/tmtx/res-sys/pkg/validator"
)

type messageBus struct {
	queue        chan Message
	callbacks    map[MessageKey][]Callback
	errorHandler ErrorHandler
}

func NewBasicMessageBus(errorHandler ErrorHandler) MessageBus {
	return &messageBus{
		make(chan Message, 5),
		map[MessageKey][]Callback{},
		errorHandler,
	}
}

func (b *messageBus) Dispatch(m Message) {
	b.queue <- m
}

func (b *messageBus) DispatchSync(m Message) (validator.Messages, error) {
	if b.callbacks[m.Key] == nil {
		return nil, fmt.Errorf("No callbacks registered for key: " + string(m.Key))
	}

	return b.executeCallbacks(m)
}

func (b *messageBus) Subscribe(key MessageKey, cb Callback) {
	b.callbacks[key] = append(b.callbacks[key], cb)
}

func (b *messageBus) Listen() {
	go func() {
		for m := range b.queue {
			if b.callbacks[m.Key] == nil {
				continue
			}
			b.executeCallbacks(m)
		}
	}()
}

func (b *messageBus) executeCallbacks(m Message) (validator.Messages, error) {
	allMessages := validator.Messages{}
	var err error
	for _, cb := range b.callbacks[m.Key] {
		validatorMessages, err := cb(m.Params)
		if err != nil {
			break
		}
		allMessages = validator.MergeMessages(allMessages, validatorMessages)
		b.errorHandler.Handle(err, validatorMessages)
	}

	return allMessages, err
}
