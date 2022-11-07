package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	connection *redis.Client
}

func NewStorage(conn *redis.Client) Storage {
	return Storage{
		connection: conn,
	}
}

func (s *Storage) Put(key, value string) {
	ctx := context.Background()

	s.connection.Set(ctx, key, value, 0)
}

func (s *Storage) Get(key string) string {
	ctx := context.Background()

	v := s.connection.Get(ctx, key)

	return v.String()
}