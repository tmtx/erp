package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/tmtx/erp/app"
	"github.com/tmtx/erp/app/server"
	"github.com/tmtx/erp/app/services/guests"
	"github.com/tmtx/erp/app/services/reservations"
	"github.com/tmtx/erp/app/services/spaces"
	"github.com/tmtx/erp/app/services/users"
	"github.com/tmtx/erp/pkg/mongo/event"
	redisbus "github.com/tmtx/erp/pkg/redis/bus"
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
	userService := users.New(basicService)
	spacesService := spaces.New(basicService)
	reservationsService := reservations.New(basicService)

	s := server.New([]server.Router{
		guestService.NewRouter(),
		userService.NewRouter(),
		spacesService.NewRouter(),
		reservationsService.NewRouter(),
	})

	app.RegisterCommandSubscribers([]app.CommandSubscriber{
		&guestService,
		&userService,
		&reservationsService,
	})

	go commandBus.Listen()

	s.Start(":8080")
}

func (eh *errorHandler) Handle(err error) {
	fmt.Println(err)
}
