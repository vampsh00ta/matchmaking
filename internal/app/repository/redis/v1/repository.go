package v1

import "context"

type Repository interface {
	AddUserToQueue(ctx context.Context, rating, tgID int) error
	DeleteUserFromQueue(ctx context.Context, rating, tgID int) error
	DeleteUsersFromQueue(ctx context.Context, rating int, tgID ...int) error
	GetIDsFromQueue(ctx context.Context, tgID int) ([]int, error)
}
