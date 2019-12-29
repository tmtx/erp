package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/app/services/guests"
	"github.com/tmtx/res-sys/pkg/mongo/event"
	redisbus "github.com/tmtx/res-sys/pkg/redis/bus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type errorHandler struct{}

func main() {
	redisBusOptions := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	commandBus, err := redisbus.NewRedisMessageBus(&redisBusOptions)
	if err != nil {
		panic(err)
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	eventRepository, err := event.NewRepository(clientOptions, "erp")
	if err != nil {
		panic(err)
	}

	basicService := app.NewBasicService(commandBus, eventRepository)
	guestService := guests.New(basicService)

	s := server.New([]server.Router{
		guestService.NewRouter(),
	})

	app.RegisterCommandSubscribers([]app.CommandSubscriber{
		&guestService,
	})

	go commandBus.Listen()

	s.Start(":8081")
}

func (eh *errorHandler) Handle(err error) {
	fmt.Println(err)
}
