package test

import (
	"context"
	"errors"
	"testing"

	"github.com/amirhnajafiz/testcontainers/pkg/storage/redis"
)

func TestRedisContainer(t *testing.T) {
	ctx := context.Background()

	redisC, err := redis.CreateRedisContainer()
	if err != nil {
		t.Error(err)

		return
	}

	redisConnection, err := redisC.Endpoint(ctx, "")
	if err != nil {
		t.Error(err)

		return
	}

	rc := redis.NewStorage(redisConnection)

	rc.Put("my-key", "my-value")

	if rc.Get("my-key") != "my-value" {
		t.Error(errors.New("redis container operation failed"))
	}
}
