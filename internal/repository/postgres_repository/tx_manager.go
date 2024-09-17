package postgresrep

import (
	"context"
	"matchmaking/pkg/client"
)

type PgxTxManager struct {
	client client.Client
	ctxm   client.CtxManager
}

func NewPgxTxManager(c client.Client) client.Manager {
	return &PgxTxManager{
		client: c,
		ctxm:   client.NewPgxCtxManager(c),
	}
}
func (m PgxTxManager) Do(ctx context.Context, f client.TxFunc) error {
	tx, err := m.client.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Commit(ctx)
	}()
	//создаем новый contex
	ctx = m.ctxm.SetDefault(ctx, tx)
	// вызываем фукнцию с новым контекстом
	err = f(ctx)
	//закрывыем транзанцию
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}
	return nil
}
