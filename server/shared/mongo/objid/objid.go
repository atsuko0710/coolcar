package objid

import (
	"coolcar/server/shared/id"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

func MustFromID(id fmt.Stringer) primitive.ObjectID {
	oid, err := FromID(id)
	if err != nil {
		panic(err)
	}
	return oid
}

// ToAccountId 转换 ObjectID to AccountID
func ToAccountId(oid primitive.ObjectID) id.AccountID {
	return id.AccountID(oid.Hex())
}

func ToTripId(oid primitive.ObjectID) id.TripID {
	return id.TripID(oid.Hex())
}