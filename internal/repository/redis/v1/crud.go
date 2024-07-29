package mongodb

import (
	"context"
	"errors"
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

func (d db) AddUserToQueue(ctx context.Context, rating, tgID int) error {
	ids, err := d.GetIDsFromQueue(ctx, rating)
	if err != nil {
		return err
	}
	for _, id := range ids {
		if id == tgID {
			return nil
		}
	}
	ids = append(ids, tgID)
	idsStr := d.encodeIntList(ids)
	if err := d.client.HSet(ctx, queueKey, rating, idsStr).Err(); err != nil {
		return err
	}
	return nil

}

func (d db) DeleteUserFromQueue(ctx context.Context, matchmaking, tgID int) error {
	currIDs, err := d.GetIDsFromQueue(ctx, matchmaking)
	if err != nil {
		return err
	}
	newIDs := make([]int, 0)
	for i := range currIDs {
		if currIDs[i] != tgID {
			newIDs = append(newIDs, currIDs[i])
		}
	}
	newIDsStr := d.encodeIntList(newIDs)
	if err := d.client.HSet(ctx, queueKey, matchmaking, newIDsStr).Err(); err != nil {
		return err
	}
	return nil
}

func (d db) DeleteUsersFromQueue(ctx context.Context, matchmaking int, tgIDs ...int) error {
	if len(tgIDs) == 0 {
		return errors.New("must provide at least one tgID")
	}
	tgIDsToDelete := make(map[int]bool)
	for _, tgID := range tgIDs {
		tgIDsToDelete[tgID] = true
	}
	currTgIDS, err := d.GetIDsFromQueue(ctx, matchmaking)
	if err != nil {
		return err
	}
	newIDs := make([]int, len(tgIDsToDelete)-len(tgIDs))
	for i := range currTgIDS {
		tgID := currTgIDS[i]
		if !tgIDsToDelete[tgID] {
			newIDs = append(newIDs, currTgIDS[i])
		}
	}
	idsStr := d.encodeIntList(newIDs)
	if err := d.client.HSet(ctx, queueKey, matchmaking, idsStr).Err(); err != nil {
		return err
	}
	return nil
}

func (d db) GetIDsFromQueue(ctx context.Context, matchmaking int) ([]int, error) {
	strRating := strconv.Itoa(matchmaking)
	res, err := d.client.HGet(ctx, queueKey, strRating).Result()
	if res == "" {
		return []int{}, nil
	}
	if err != nil {
		return nil, err
	}

	ids, err := d.decodeStrList(res)
	if err != nil {
		return []int{}, err
	}
	return ids, nil
}
