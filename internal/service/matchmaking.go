package service

import (
	"context"
	"errors"
	ipsql "matchmaking/internal/app/repository/psql"
	iredis "matchmaking/internal/app/repository/redis"
	isrvc "matchmaking/internal/app/service"
	"matchmaking/internal/entity"
	"matchmaking/pkg/client"
)

//type Service interface {
//	FindMatch(ctx context.Context,tgID int) error
//	MatchResult(ctx context.Context, tgIDWinner,tgIDLoser int) error
//}

const (
	okDiff        = 500
	baseMmrChange = 50
)

type match struct {
	rating ipsql.Rating
	txm    client.Manager
	queue  iredis.Queue
}

func NewMatch(rating ipsql.Rating, txm client.Manager, queue iredis.Queue) isrvc.Match {
	return &match{
		rating: rating,
		txm:    txm,
		queue:  queue,
	}
}

func (s match) Find(ctx context.Context, tgID int) (int, error) {
	res, err := s.rating.GetByTgID(ctx, tgID)
	if err != nil {
		return -1, err
	}
	user := entity.User{
		Rating: res,
		TgID:   tgID,
	}
	currUsers, err := s.queue.GetUsers(ctx)
	if err != nil {
		return -1, nil
	}

	foundUser := findLeastDiff(user, currUsers...)
	if foundUser.TgID == tgID {
		return -1, errors.New("already searching fight")
	}
	if foundUser.TgID == -1 {
		if err := s.queue.AddUser(ctx, user); err != nil {
			return -1, err
		}
		return -1, nil
	}

	if err := s.queue.DeleteUserByTgID(ctx, foundUser.TgID); err != nil {
		return -1, nil
	}
	return foundUser.TgID, nil
}

func (s match) Result(ctx context.Context, tgIDWinner, tgIDLoser int) error {
	if tgIDWinner == tgIDLoser {
		return errors.New("values are equal")
	}
	//tx, err := s.ra.Begin(ctx)

	var winnerRating, loserRating int
	err := s.txm.Do(ctx, func(ctx context.Context) error {
		var err error

		winnerRating, err = s.rating.GetByTgID(ctx, tgIDWinner)
		if err != nil {
			return err
		}
		loserRating, err = s.rating.GetByTgID(ctx, tgIDLoser)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	resultChangeMmr := calculateRatingChange(winnerRating, loserRating)

	//add transactions
	//winner
	err = s.txm.Do(ctx, func(ctx context.Context) error {
		if err := s.rating.UpdateByTgID(ctx, tgIDWinner, resultChangeMmr); err != nil {
			return err
		}
		if err := s.rating.UpdateByTgID(ctx, tgIDLoser, -resultChangeMmr); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
