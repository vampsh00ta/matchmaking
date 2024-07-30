package postgresrep

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"matchmaking/pkg/client"
)

//	type Tx interface {
//		GetRating(ctx context.Context, tgID int) (int, error)
//		UpdateRating(ctx context.Context, tgID int, matchmakingChange int) error
//		InitRating(ctx context.Context, tgID, matchmaking int) error
//		getDB(ctx context.Context) client.Client
//	}
type (
	txKey     struct{}
	clientKey struct{}
)

func (db pg) Rollback(ctx context.Context) error {
	txCtx := ctx.Value(txKey{})
	tx, ok := txCtx.(pgx.Tx)
	if !ok {
		return fmt.Errorf("no tx")
	}

	if err := tx.Rollback(ctx); err != nil {
		return err
	}
	return nil
}

func (db pg) Commit(ctx context.Context) error {
	txCtx := ctx.Value(txKey{})
	tx, ok := txCtx.(pgx.Tx)
	if !ok {
		return fmt.Errorf("no tx")
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (db pg) Begin(ctx context.Context) (context.Context, error) {
	tx, err := db.client.Begin(ctx)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, txKey{}, tx)
	ctx = context.WithValue(ctx, clientKey{}, client.Client(tx))

	return ctx, nil
}

func (db pg) getDB(ctx context.Context) client.Client {
	tx := ctx.Value(clientKey{})
	txModel, ok := tx.(client.Client)
	if !ok {
		txModel = db.client
	}
	return txModel
}
