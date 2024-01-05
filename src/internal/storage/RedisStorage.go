package storage

import (
	"github.com/Alextek777/fily/src/internal/config"
)

type RedisStore struct {
	storage string
}

func NewRedisStore(cfg *config.Config) (*RedisStore, error) {
	return &RedisStore{storage: ""}, nil
}

func (s *RedisStore) InitStorage() error {
	return nil
}
