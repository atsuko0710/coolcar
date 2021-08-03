package dao

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoUri string

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

func TestMain(m *testing.M)  {
	c, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	resp, err := c.ContainerCreate(ctx, &container.Config{
		Image: "mongo:4.1",
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "27018",
				},
			},
		},
	}, &network.NetworkingConfig{}, nil, "")
	if err != nil {
		panic(err)
	}
	defer func ()  {
		err := c.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{
			Force: true,
		})
		log.Fatalf("error removing container:%v", err)
	}()

	err = c.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	
	fmt.Println("container started")

	inspRes, err := c.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}

	hostPort := inspRes.NetworkSettings.Ports["27017/tcp"][0]
	mongoUri = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)

	err = c.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}