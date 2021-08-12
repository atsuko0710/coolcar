package dao

import (
	"context"
	"os"
	"testing"

	rentalpb "coolcar/server/rental/api/gen/v1"

	mongotesting "coolcar/server/shared/mongo/testing"
)

var mongoURI string

func TestCreateTrip(t *testing.T) {

	mongoURI = "mongodb://root:123456@127.0.0.1:27017"

	c := context.Background()

	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	db := mc.Database("coolcar")
	err = mongotesting.SetupIndexes(c, db)
	if err != nil {
		t.Fatalf("cannot setup index: %v", err)
	}

	m := NewMongo(db)
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "startpoint",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName: "endpoint",
			FeeCent: 10000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 112,
			},
		},
		Status: rentalpb.TripStatus_IN_PROGRESS,
	})
	if err != nil {
		t.Errorf("cannot create trip %v", err)
	}

	t.Errorf("%+v", tr)
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
