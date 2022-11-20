package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	// nats docker image information
	natsImageName = "nats"
	natsImageTag  = "latest"

	// nats testing parameters
	natsTopic = "some-private-key"
	natsValue = "55#ou8aAApo#e9kkd"
)

// createRedisContainer
// generates a new nats' container.
func createNatsContainer(ctx context.Context) (testcontainers.Container, error) {
	// container request
	req := testcontainers.ContainerRequest{
		Image:        natsImageName + ":" + natsImageTag,
		ExposedPorts: []string{"4222/tcp"},
		WaitingFor:   wait.ForLog("Listening for client connections"),
	}

	// creating a new redis container
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

// TestNatsContainer
// testing nats container.
func TestNatsContainer(t *testing.T) {
	// creating a new context
	ctx := context.Background()

	// creating nats container
	natsContainer, err := createNatsContainer(ctx)
	if err != nil {
		t.Error(fmt.Errorf("create nats container failed:\n\t%v\n", err))

		return
	}

	// getting nats connection
	natsConnection, err := natsContainer.Endpoint(ctx, "")
	if err != nil {
		t.Error(fmt.Errorf("getting nats connection failed:\n\t%v\n", err))

		return
	}

	// testing nats connection
	{
		// opening connection
		nc, er := nats.Connect(natsConnection)
		if er != nil {
			t.Error(fmt.Errorf("connecting to nats container failed:\n\t%v\n", er))

			return
		}

		// async subscriber
		go func() {
			_, e := nc.Subscribe(natsTopic, func(m *nats.Msg) {
				log.Printf("Received a message:\n\t%s\n", string(m.Data))
			})
			if e != nil {
				t.Error(fmt.Errorf("subscribe over topic failed:\n\t%v\n", e))
			}
		}()

		// publish over topic
		if e := nc.Publish(natsTopic, []byte(natsValue)); e != nil {
			t.Error(fmt.Errorf("publish over topic failed:\n\t%v\n", e))
		}
	}
}
