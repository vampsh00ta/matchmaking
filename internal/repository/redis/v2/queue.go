package v2

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	iredis "matchmaking/internal/app/repository/redis/v2"
	"matchmaking/internal/entity"
	"strconv"
)

type queue struct {
	client *redis.Client
}

func NewQueue(client *redis.Client) iredis.Queue {
	return &queue{client: client}
}

const (
	queueKey     = "search-queue"
	ratingGroup  = 100
	possibleDiff = 500
	separator    = ";"
)

func (d queue) AddUser(ctx context.Context, userToAdd entity.User) error {
	users, err := d.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.TgID == userToAdd.TgID {
			return errors.New("already in queue")
		}
	}
	if err := d.client.HSet(ctx, queueKey, userToAdd.TgID, userToAdd.Rating).Err(); err != nil {
		return err
	}
	return nil

}

func (d queue) DeleteUserByTgID(ctx context.Context, tgID int) error {
	tgIDStr := strconv.Itoa(tgID)
	if err := d.client.HDel(ctx, queueKey, tgIDStr).Err(); err != nil {
		return err
	}
	return nil
}

func (d queue) DeleteUsersByTgIDs(ctx context.Context, tgIDs ...int) error {

	tgIDStrs := make([]string, len(tgIDs))
	for i, tgID := range tgIDs {
		tgIDStrs[i] = strconv.Itoa(tgID)
	}
	if err := d.client.HDel(ctx, queueKey, tgIDStrs...).Err(); err != nil {
		return err
	}
	return nil
}

func (d queue) GetUsers(ctx context.Context) ([]entity.User, error) {
	m, err := d.client.HGetAll(ctx, queueKey).Result()
	if err != nil {
		return nil, err
	}
	res := make([]entity.User, len(m))
	var i int
	for key, value := range m {
		tgID, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		rating, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		res[i] = entity.User{
			TgID:   tgID,
			Rating: rating,
		}
		i++
	}
	return res, nil
}
