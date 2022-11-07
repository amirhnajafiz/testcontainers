package redis

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	imageName = "redis"
	imageTag  = "latest"
)

func CreateRedisContainer() (testcontainers.Container, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        imageName + ":" + imageTag,
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
