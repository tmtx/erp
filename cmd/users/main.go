package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/tmtx/res-sys/app"
	"github.com/tmtx/res-sys/app/server"
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

	userService := users.New(basicService)

	s := server.New(
		[]server.Router{
			userService.NewRouter(),
		},
		os.Getenv("COOKIE_STORE_SECRET"),
		os.Getenv("DEFAULT_CORS_ORIGIN"),
	)

	go commandBus.Listen()

	s.Start(os.Getenv("ECHO_HOST"))
}

func (eh *errorHandler) Handle(err error) {
	fmt.Println(err)
}
