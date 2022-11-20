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

type natsContainer struct {
	testcontainers.Container
	URI string
}

// setupNats
// generates a new nats cluster container.
func setupNats(ctx context.Context) (*natsContainer, error) {
	// container build request
	req := testcontainers.ContainerRequest{
		Image:        "nats:latest",
		ExposedPorts: []string{"4222/tcp"},
		WaitingFor:   wait.ForLog("Listening for client connections"),
	}

	// building a generic container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("building generic container failed: %v", err)
	}

	// mapped port
	mappedPort, err := container.MappedPort(ctx, "4222")
	if err != nil {
		return nil, fmt.Errorf("getting mapped port failed: %v", err)
	}

	// getting the container ip
	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting host ip failed: %v", err)
	}

	// generating the nats cluster uri
	uri := fmt.Sprintf("nats://%s:%s", hostIP, mappedPort.Port())

	// creating a new nats container
	return &natsContainer{Container: container, URI: uri}, nil
}

// TestNatsContainer
// testing nats container.
func TestNatsContainer(t *testing.T) {
	// checking test flag
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// nats cluster testing parameters
	const (
		natsTopic = "some-private-key"
		natsValue = "55#ou8aApo#e9kkd"
	)

	// creating a new context
	ctx := context.Background()

	// creating nats container
	container, err := setupNats(ctx)
	if err != nil {
		t.Error(fmt.Errorf("create nats container failed:\n\t%v\n", err))

		return
	}

	// cleaning container after test is complete
	t.Cleanup(func() {
		t.Log("terminating container")

		if er := container.Terminate(ctx); er != nil {
			t.Errorf("failed to terminate container: :%v", er)
		}
	})

	// testing Nats connection.
	// You will likely want to wrap your Nats package of choice in an
	// interface to aid in unit testing and limit lock-in throughout your
	// codebase but that's out of scope for this example
	{
		// opening connection
		nc, er := nats.Connect(container.URI)
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
