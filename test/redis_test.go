package test

import (
	"context"
	"errors"
	"testing"

	"github.com/amirhnajafiz/testcontainers/pkg/storage/redis"
)

// TestRedisContainer
// creating a redis container by building one and
// send requests to redis connections.
func TestRedisContainer(t *testing.T) {
	ctx := context.Background()

	// creating a new container.
	redisC, err := redis.CreateRedisContainer()
	if err != nil {
		t.Error(err)

		return
	}

	// get container connection.
	redisConnection, err := redisC.Endpoint(ctx, "")
	if err != nil {
		t.Error(err)

		return
	}

	// creating a new storage
	rc := redis.NewStorage(redisConnection)

	// testing storage
	rc.Put("my-key", "my-value")

	if rc.Get("my-key") != "my-value" {
		t.Error(errors.New("redis container operation failed"))
	}
}
