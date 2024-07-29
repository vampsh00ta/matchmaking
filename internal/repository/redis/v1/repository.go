package mongodb

import (
	"github.com/redis/go-redis/v9"
	iredis "matchmaking/internal/app/repository/redis/v1"
)

type db struct {
	client *redis.Client
}

func New(client *redis.Client) iredis.Repository {
	return db{client: client}
}
