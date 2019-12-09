package bus

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/validator"
)

type redisBus struct {
	client        *redis.Client
	subscriptions map[bus.MessageKey][]bus.Callback
}

type MessageBus interface {
	Dispatch(m bus.Message)
	DispatchSync(m bus.Message)
	Subscribe(key bus.MessageKey, cb bus.Callback)
	Listen()
}

func NewRedisMessageBus(options *redis.Options) (bus.MessageBus, error) {
	client := redis.NewClient(options)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return redisBus{
		client:        client,
		subscriptions: map[bus.MessageKey][]bus.Callback{},
	}, nil
}

func (b redisBus) Dispatch(m bus.Message) {
	b.client.Publish(string(m.Key), &m)
}

func (b redisBus) DispatchSync(m bus.Message) (validator.Messages, error) {
	if b.subscriptions[m.Key] == nil {
		return nil, fmt.Errorf("No callbacks registered for key: " + string(m.Key))
	}

	return b.executeCallbacks(m)
}

func (b redisBus) Subscribe(key bus.MessageKey, cb bus.Callback) {
	b.subscriptions[key] = append(b.subscriptions[key], cb)
}

func (b redisBus) Listen() {
	for key := range b.subscriptions {
		go b.handleSubscription(key)
	}
}

func (b redisBus) handleSubscription(key bus.MessageKey) {
	pubsub := b.client.Subscribe(string(key))

	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	for msg := range pubsub.Channel() {
		var m bus.Message
		m.UnmarshalBinary([]byte(msg.Payload))
		b.executeCallbacks(m)
	}
}

func (b redisBus) executeCallbacks(m bus.Message) (validator.Messages, error) {
	var allMessages validator.Messages
	var err error
	var validatorMessages validator.Messages
	for _, cb := range b.subscriptions[m.Key] {
		validatorMessages, err = cb(m.Params)
		if err != nil {
			break
		}
		allMessages = validator.MergeMessages(allMessages, validatorMessages)
	}

	return allMessages, err
}
