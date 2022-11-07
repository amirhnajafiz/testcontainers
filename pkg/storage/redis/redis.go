package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Storage
// manages the redis connections.
type Storage struct {
	connection *redis.Client
}

// NewStorage
// creates a new storage struct.
func NewStorage(conn string) Storage {
	// opening a new redis connection.
	client := redis.NewClient(&redis.Options{
		Addr: conn,
	})

	return Storage{
		connection: client,
	}
}

// Put
// puts a new set of key value in redis database.
func (s *Storage) Put(key, value string) {
	ctx := context.Background()

	s.connection.Set(ctx, key, value, 0)
}

// Get
// restore a value by key from redis database.
func (s *Storage) Get(key string) string {
	ctx := context.Background()

	v := s.connection.Get(ctx, key)

	return v.String()
}
