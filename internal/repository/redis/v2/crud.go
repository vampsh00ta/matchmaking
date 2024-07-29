package v2

import (
	"context"
	"errors"
	"matchmaking/internal/entity"
	"strconv"
)

//type Repository interface {
//	AddUserToQueue(ctx context.Context, tgID, matchmaking int) error
//	DeleteUserFromQueue(ctx context.Context, matchmaking,tgID int) error
//	DeleteUsersFromQueue(ctx context.Context, matchmaking int, tgID ...int) error
//	GetIDsFromQueue(ctx context.Context, tgID int) ([]int, error)
//}

// map[key:value] structure
//
//	queue : {
//		0: [tgs...],
//	 100: [tgs...]
//	 200: [tgs...]

// ..
// ..
// 10000: [tgs...]
// }
const (
	queueKey     = "search-queue"
	ratingGroup  = 100
	possibleDiff = 500
	separator    = ";"
)

func (d db) AddUserToQueue(ctx context.Context, userToAdd entity.User) error {
	users, err := d.GetUsersInQueue(ctx)
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

func (d db) DeleteUserFromQueue(ctx context.Context, tgID int) error {
	tgIDStr := strconv.Itoa(tgID)
	if err := d.client.HDel(ctx, queueKey, tgIDStr).Err(); err != nil {
		return err
	}
	return nil
}

func (d db) DeleteUsersFromQueue(ctx context.Context, tgIDs ...int) error {

	tgIDStrs := make([]string, len(tgIDs))
	for i, tgID := range tgIDs {
		tgIDStrs[i] = strconv.Itoa(tgID)
	}
	if err := d.client.HDel(ctx, queueKey, tgIDStrs...).Err(); err != nil {
		return err
	}
	return nil
}

func (d db) GetUsersInQueue(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	if err := d.client.HGetAll(ctx, queueKey).Scan(&users); err != nil {
		return nil, err
	}

	return users, nil
}
