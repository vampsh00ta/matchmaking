package postgresrep

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
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

func (db pg) InitRating(ctx context.Context, tgID, rating int) error {
	q := `insert into user_rating(tg_id,rating) values($1,$2) returning id`
	client := db.getDB(ctx)

	if err := client.QueryRow(ctx, q, tgID, rating).Scan(&tgID); err != nil {
		return err
	}
	return nil
}

func (db pg) GetRating(ctx context.Context, tgID int) (int, error) {
	var rating int
	q := `select  rating from   user_rating where tg_id=$1`

	client := db.getDB(ctx)

	err := client.QueryRow(ctx, q, tgID).Scan(&rating)

	if errors.Is(err, pgx.ErrNoRows) {
		if err := db.InitRating(ctx, tgID, initRating); err != nil {
			return -1, err
		}
		return initRating, nil
	}
	if err != nil {

		return -1, err
	}

	return rating, nil
}

func (db pg) UpdateRating(ctx context.Context, tgID int, ratingChange int) error {
	q := `update user_rating set rating=rating + $2 where tg_id=$1 returning id`

	client := db.getDB(ctx)

	if err := client.QueryRow(ctx, q, tgID, ratingChange).Scan(&ratingChange); err != nil {
		return err
	}
	return nil
}
