package redis

import (
	"context"
	"matchmaking/internal/entity"
)

type Queue interface {
	AddUser(ctx context.Context, userToAdd entity.User) error
	DeleteUserByTgID(ctx context.Context, tgID int) error
	DeleteUsersByTgIDs(ctx context.Context, tgIDs ...int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
}
