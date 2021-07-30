package dao

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountId(t *testing.T)  {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://root:123456@127.0.0.1:27017/coolcar?authSource=admin&readPreference=primary&appname=mongodb-vscode%200.6.10&directConnection=true&ssl=false"))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	defer mc.Disconnect(c)

	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountId(c, "abc")
	if err != nil {
		t.Fatalf("faild resolve account id for abc:%v", err)
	} else {
		want := "61035ee3d2eb3acb17da3c25"
		if want != id {
			t.Fatalf("resolve account id:want: %q got: %q", want, id)
		}
	}
}