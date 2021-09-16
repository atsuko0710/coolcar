package dao

import (
	"context"
	"coolcar/server/shared/id"
	mgutil "coolcar/server/shared/mongo"
	"coolcar/server/shared/mongo/objid"
	"fmt"

	carpb "coolcar/server/car/api/gen/v1"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

type CarRecord struct {
	mgutil.IDField `bson:"inline"`
	Car            *carpb.Car `bson:"car"`
}

func (m *Mongo) CreateCar(c context.Context) (*CarRecord, error) {
	r := &CarRecord{
		Car: &carpb.Car{
			Position: &carpb.Location{
				Latitude: 30,
				Longitude: 120,
			},
			Status: carpb.CarStatus_LOCKED,
		},
	}
	r.ID = mgutil.NewObjID()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (m *Mongo) GetCar(c context.Context, id id.CarID) (*CarRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id %v:", err)
	}
	res := m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
	})
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("cannot find one %v:", err)
	}

	var cr CarRecord
	err = res.Decode(&cr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode result:%v", err)
	}
	return &cr, nil
}

func (m *Mongo) GetCars(c context.Context) ([]*CarRecord, error) {
	filter := bson.M{}
	res, err := m.col.Find(c, filter, options.Find())
	if err != nil {
		return nil, err
	}

	var cars []*CarRecord
	for res.Next(c) {
		var car CarRecord
		err := res.Decode(&car)
		if err != nil {
			return nil, err
		}
		cars = append(cars, &car)
	}
	return cars, nil
}