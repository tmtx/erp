package event

import (
	"context"
	"time"

	"github.com/tmtx/erp/pkg/bus"
	"github.com/tmtx/erp/pkg/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client     *mongo.Client
	dbName     string
	eventTypes map[bus.MessageKey]interface{}
}

func NewRepository(options *options.ClientOptions, dbName string) (r event.Repository, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options)
	if err == nil {
		err = client.Ping(ctx, readpref.Primary())
	}

	return &mongoRepository{
		client,
		dbName,
		map[bus.MessageKey]interface{}{},
	}, err
}

func (r *mongoRepository) Store(ctx context.Context, e event.Event) (err error) {
	c := r.client.Database(r.dbName).Collection("events")

	b, err := bson.Marshal(bson.D{
		{"key", e.Key},
		{"params", e.Params},
	})
	if err != nil {
		return err
	}

	_, err = c.InsertOne(ctx, b)

	return err
}

func (r *mongoRepository) FindOneWithFilter(
	ctx context.Context,
	filter event.Filter,
) (event.Event, error) {
	c := r.client.Database(r.dbName).Collection("events")

	e := event.Event{}
	err := c.FindOne(ctx, filter).Decode(&e)

	return e, err
}

func (r *mongoRepository) FindAllWithFilter(
	ctx context.Context,
	filter event.Filter,
) ([]event.Event, error) {
	c := r.client.Database(r.dbName).Collection("events")

	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := []event.Event{}
	for cursor.Next(ctx) {
		e := event.Event{}
		err = cursor.Decode(&e)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	return result, err
}
