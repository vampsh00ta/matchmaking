package service

import (
	"context"
	"errors"
	"matchmaking/internal/entity"
	"math"
)

//type Service interface {
//	FindMatch(ctx context.Context,tgID int) error
//	MatchResult(ctx context.Context, tgIDWinner,tgIDLoser int) error
//}

const (
	okDiff = 500
)

func (s service) FindMatch(ctx context.Context, tgID int) (int, error) {
	rating, err := s.psql.GetRating(ctx, tgID)
	if err != nil {
		return -1, err
	}
	user := entity.User{
		Rating: rating,
		TgID:   tgID,
	}
	currUsers, err := s.redis.GetUsersInQueue(ctx)
	if err != nil {
		return -1, nil
	}

	foundUser := s.findLeastDiff(user, currUsers...)
	if foundUser.TgID == tgID {
		return -1, errors.New("already searching fight")
	}
	if foundUser.TgID == -1 {
		if err := s.redis.AddUserToQueue(ctx, user); err != nil {
			return -1, err
		}
		return -1, nil
	}

	if err := s.redis.DeleteUserFromQueue(ctx, foundUser.TgID); err != nil {
		return -1, nil
	}
	return foundUser.TgID, nil
}

func (s service) findLeastDiff(newUser entity.User, users ...entity.User) entity.User {
	var res entity.User
	res.TgID = -1
	currDiff := okDiff

	for _, user := range users {
		floatDiff := float64(newUser.Rating - user.Rating)
		if int(math.Abs(floatDiff)) <= currDiff {
			currDiff = newUser.Rating - user.Rating
			res = user
		}
	}
	return res
}

func (s service) MatchResult(ctx context.Context, tgIDWinner, tgIDLoser int) error {
	//TODO implement me
	panic("implement me")
}
