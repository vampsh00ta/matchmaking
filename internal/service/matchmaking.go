package service

import (
	"context"
	"errors"
)

//type Service interface {
//	FindMatch(ctx context.Context,tgID int) error
//	MatchResult(ctx context.Context, tgIDWinner,tgIDLoser int) error
//}

func (s service) FindMatch(ctx context.Context, tgID int) (int, error) {
	rating, err := s.psql.GetRating(ctx, tgID)
	if err != nil {
		return -1, err
	}
	currIDs, err := s.redis.GetIDsFromQueue(ctx, rating)
	if err != nil {
		return -1, nil
	}
	if len(currIDs) == 0 {
		if err := s.redis.AddUserToQueue(ctx, rating, tgID); err != nil {
			return -1, err
		}
		return -1, nil
	}
	id := s.findLeastDiff(tgID, currIDs...)
	if id == tgID {
		return -1, errors.New("already searching fight")
	}
	if err := s.redis.DeleteUserFromQueue(ctx, rating, id); err != nil {
		return -1, nil
	}
	return id, nil
}

func (s service) findLeastDiff(searchingTgID int, currIDs ...int) int {
	var res int
	currDiff := searchingTgID

	for _, id := range currIDs {
		if (searchingTgID - id) < currDiff {
			currDiff = searchingTgID - id
			res = id
		}
	}
	return res
}

func (s service) MatchResult(ctx context.Context, tgIDWinner, tgIDLoser int) error {
	//TODO implement me
	panic("implement me")
}
