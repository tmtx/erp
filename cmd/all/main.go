package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
	"github.com/tmtx/res-sys/app/services/guests"
	"github.com/tmtx/res-sys/app/services/reservations"
	"github.com/tmtx/res-sys/app/services/spaces"
	"github.com/tmtx/res-sys/app/services/users"
	"github.com/tmtx/res-sys/pkg/env"
	"github.com/tmtx/res-sys/pkg/mongo/event"
	redisbus "github.com/tmtx/res-sys/pkg/redis/bus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type errorHandler struct{}

func main() {
	env.LoadEnvironment()

	redisBusOptions := redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	}
	commandBus, err := redisbus.NewRedisMessageBus(&redisBusOptions)
	if err != nil {
		panic(err)
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONN"))
	eventRepository, err := event.NewRepository(clientOptions, "erp")
	if err != nil {
		panic(err)
	}

	basicService := app.NewBasicService(commandBus, eventRepository)

	guestService := guests.New(basicService)
	userService := users.New(basicService)
	spacesService := spaces.New(basicService)
	reservationsService := reservations.New(basicService)

	s := server.New(
		[]server.Router{
			guestService.NewRouter(),
			userService.NewRouter(),
			spacesService.NewRouter(),
			reservationsService.NewRouter(),
		},
		os.Getenv("COOKIE_STORE_SECRET"),
		os.Getenv("DEFAULT_CORS_ORIGIN"),
	)

	app.RegisterCommandSubscribers([]app.CommandSubscriber{
		&guestService,
		&userService,
		&reservationsService,
	})

	go commandBus.Listen()

	s.Start(os.Getenv("ECHO_HOST"))
}

func (eh *errorHandler) Handle(err error) {
	fmt.Println(err)
}
