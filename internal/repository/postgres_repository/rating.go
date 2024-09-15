package postgresrep

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	ipsql "matchmaking/internal/app/repository/psql"
	"matchmaking/pkg/client"
)

//		type Repository interface {
//			GetRating(ctx context.Context,tgID int) error
//			UpdateRating(ctx context.Context,tgID int,matchmakingChange int) error
//	     InitRating(ctx context.Context, tgID int) (int,error)
//
// }

const (
	initRating = 1000
)

type rating struct {
	db client.DbManager
}

func NewRating(c client.Client) ipsql.Rating {
	return &rating{
		db: NewPgxManager(c),
	}
}
func (rep rating) Init(ctx context.Context, tgID, rating int) error {
	q := `insert into user_rating(tg_id,rating) values($1,$2) returning id`
	client, err := rep.db.DefaultTrOrDb(ctx)
	if err != nil {
		return err
	}
	if err := client.QueryRow(ctx, q, tgID, rating).Scan(&tgID); err != nil {
		return err
	}
	return nil
}

func (rep rating) GetByTgID(ctx context.Context, tgID int) (int, error) {
	var rating int
	q := `select  rating from   user_rating where tg_id=$1`

	client, err := rep.db.DefaultTrOrDb(ctx)
	if err != nil {
		return -1, err
	}

	err = client.QueryRow(ctx, q, tgID).Scan(&rating)

	if errors.Is(err, pgx.ErrNoRows) {
		if err := rep.Init(ctx, tgID, initRating); err != nil {
			return -1, err
		}
		return initRating, nil
	}
	if err != nil {

		return -1, err
	}

	return rating, nil
}

func (rep rating) UpdateByTgID(ctx context.Context, tgID int, ratingChange int) error {
	q := `update user_rating set rating=rating + $2 where tg_id=$1 returning id`

	client, err := rep.db.DefaultTrOrDb(ctx)
	if err != nil {
		return err
	}

	if err := client.QueryRow(ctx, q, tgID, ratingChange).Scan(&ratingChange); err != nil {
		return err
	}
	return nil
}
