package service

import (
	"context"
	"errors"
	ipsql "matchmaking/internal/app/repository/psql"
	iredis "matchmaking/internal/app/repository/redis/v2"
	isrvc "matchmaking/internal/app/service"
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

type match struct {
	psql  ipsql.Repository
	queue iredis.Queue
}

func NewMatch(psql ipsql.Repository, queue iredis.Queue) isrvc.Match {
	return &match{
		psql:  psql,
		queue: queue,
	}
}

func (s match) Find(ctx context.Context, tgID int) (int, error) {
	rating, err := s.psql.GetRating(ctx, tgID)
	if err != nil {
		return -1, err
	}
	user := entity.User{
		Rating: rating,
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
	tx, err := s.psql.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = s.psql.Rollback(tx)
	}()

	winnerRating, err := s.psql.GetRating(tx, tgIDWinner)
	if err != nil {
		return err
	}
	loserRating, err := s.psql.GetRating(tx, tgIDLoser)
	if err != nil {
		return err
	}
	resultChangeMmr := calculateRatingChange(winnerRating, loserRating)

	//add transactions
	//winner
	if err := s.psql.UpdateRating(tx, tgIDWinner, resultChangeMmr); err != nil {
		return err
	}
	//loser
	if err := s.psql.UpdateRating(tx, tgIDLoser, -resultChangeMmr); err != nil {
		return err
	}

	if err := s.psql.Commit(tx); err != nil {
		return err
	}
	return nil
}
