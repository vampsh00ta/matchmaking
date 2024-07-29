package service

import (
	"context"
	"errors"
	"matchmaking/internal/entity"
)

//type Service interface {
//	FindMatch(ctx context.Context,tgID int) error
//	MatchResult(ctx context.Context, tgIDWinner,tgIDLoser int) error
//}

const (
	okDiff        = 500
	baseMmrChange = 50
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

func (s service) MatchResult(ctx context.Context, tgIDWinner, tgIDLoser int) error {
	winnerRating, err := s.psql.GetRating(ctx, tgIDWinner)
	if err != nil {
		return err
	}
	loserRating, err := s.psql.GetRating(ctx, tgIDLoser)
	if err != nil {
		return err
	}
	resultChangeMmr := s.calculateRatingChange(winnerRating, loserRating)

	//add transactions
	//winner
	if s.psql.UpdateRating(ctx, tgIDWinner, resultChangeMmr); err != nil {
		return err
	}
	//loser
	if s.psql.UpdateRating(ctx, tgIDLoser, -resultChangeMmr); err != nil {
		return err
	}
	return nil
}
