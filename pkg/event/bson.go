package event

import (
	"bytes"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func (u *UUID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.TypeString, bsoncore.AppendString(nil, u.String()), nil
}

func (u *UUID) UnmarshalBSONValue(t bsontype.Type, raw []byte) error {
	b := bytes.Trim(raw, "% \x00")
	result, err := uuid.ParseBytes(b)
	u.UUID = result

	return err
}
