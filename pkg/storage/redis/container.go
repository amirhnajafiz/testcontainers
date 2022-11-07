package redis

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	// redis image information
	imageName = "redis"
	imageTag  = "latest"
)

// CreateRedisContainer
// generates a new redis container.
func CreateRedisContainer() (testcontainers.Container, error) {
	// context
	ctx := context.Background()

	// container request
	req := testcontainers.ContainerRequest{
		Image:        imageName + ":" + imageTag,
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	// creating a new redis container
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
