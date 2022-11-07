package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/amirhnajafiz/testcontainers/pkg/storage/redis"
	redisSKD "github.com/go-redis/redis/v9"
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

	redisClient := redisSKD.NewClient(&redisSKD.Options{
		Addr: redisConnection,
	})

	redisClient.Set(ctx, "my-redis", "some value", time.Second*60)

	v := redisClient.Get(ctx, "my-redis")
	if v.String() != "some value" {
		t.Error(errors.New("redis container failed"))
	}
}
