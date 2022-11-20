package test

import (
	"context"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/testcontainers/testcontainers-go"
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
