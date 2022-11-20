package test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	// redis image information
	redisImageName = "redis"
	redisImageTag  = "latest"
)

// CreateRedisContainer
// generates a new redis container.
func createRedisContainer(ctx context.Context) (testcontainers.Container, error) {
	// container request
	req := testcontainers.ContainerRequest{
		Image:        redisImageName + ":" + redisImageTag,
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	// creating a new redis container
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

// TestRedisContainer
// creating a redis container by building one and
// send requests to redis connections.
func TestRedisContainer(t *testing.T) {
	// creating a new context
	ctx := context.Background()

	// creating a new container.
	redisC, err := createRedisContainer(ctx)
	if err != nil {
		t.Error(fmt.Errorf("create redis container failed:\n\t%v\n", err))

		return
	}

	// get container connection.
	redisConnection, err := redisC.Endpoint(ctx, "")
	if err != nil {
		t.Error(err)

		return
	}

	// opening a new redis connection.
	client := redis.NewClient(&redis.Options{
		Addr: redisConnection,
	})

	// testing storage
	client.Set(ctx, "my-key", "my-value", 0)

	// get redis pair for test
	if v := client.Get(ctx, "my-key"); v.String() != "my-value" {
		t.Error(errors.New("redis container operation failed"))
	}
}
