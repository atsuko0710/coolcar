package dao

import (
	"fmt"

	mgutil "coolcar/server/shared/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Mongo struct {
	col *mongo.Collection
}

const openIDField = "open_id"

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) ResolveAccountId(c context.Context, OpenID string) (string, error) {
	insertedID := mgutil.NewObjID()

	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: OpenID,
	}, mgutil.SetOnInsert(bson.M{
		mgutil.IDFieldName: insertedID,
		openIDField:        OpenID,
	}), options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))
		
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot FindOneAndUpdate:%v", err)
	}

	var row mgutil.IDField

	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result:%v", err)
	}
	return row.ID.Hex(), nil
}
