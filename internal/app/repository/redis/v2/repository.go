package v2

import (
	"context"
	"matchmaking/internal/entity"
)

type Repository interface {
	AddUserToQueue(ctx context.Context, userToAdd entity.User) error
	DeleteUserFromQueue(ctx context.Context, tgID int) error
	DeleteUsersFromQueue(ctx context.Context, tgIDs ...int) error
	GetUsersInQueue(ctx context.Context) ([]entity.User, error)
}

type Queue interface {
	AddUser(ctx context.Context, userToAdd entity.User) error
	DeleteUserByTgID(ctx context.Context, tgID int) error
	DeleteUsersByTgIDs(ctx context.Context, tgIDs ...int) error
	GetUsers(ctx context.Context) ([]entity.User, error)
}
