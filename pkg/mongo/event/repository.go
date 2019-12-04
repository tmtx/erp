package event

import (
	"context"
	"fmt"
	"time"

	"github.com/tmtx/erp/pkg/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client *mongo.Client
	dbName string
}

func NewRepository(options *options.ClientOptions, dbName string) (r event.Repository, err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options)
	if err == nil {
		err = client.Ping(ctx, readpref.Primary())
	}

	return &mongoRepository{client, dbName}, err
}

func (r *mongoRepository) Store(ctx context.Context, e event.Event) (err error) {
	c := r.client.Database(r.dbName).Collection("events")

	b, err := bson.Marshal(bson.D{
		{"key", e.Key},
		{"params", e.Params},
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(err)

	_, err = c.InsertOne(ctx, b)
	fmt.Println(err)

	return err
}

func (r *mongoRepository) FindOneWithFilter(
	ctx context.Context,
	filter event.Filter,
	result *event.Event,
) error {
	c := r.client.Database(r.dbName).Collection("events")

	fmt.Println(filter)
	return c.FindOne(ctx, filter).Decode(&result)
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

	return result, nil
}
