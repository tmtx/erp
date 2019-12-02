package app

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func (u UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.TypeString, bsoncore.AppendString(nil, u.String()), nil
}

func (u *UUID) UnmarshalBSONValue(t bsontype.Type, raw []byte) error {
	result, err := uuid.FromBytes(raw)
	u = &UUID{result}

	return err
}

func NewUUID() UUID {
	uuid := uuid.New()
	return UUID{uuid}
}
