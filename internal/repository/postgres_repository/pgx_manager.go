package postgresrep

import (
	"context"
	"matchmaking/pkg/client"
)

type PgxManager struct {
	db   client.Client
	ctxm client.CtxManager
}

func NewPgxManager(db client.Client) client.DbManager {
	return &PgxManager{
		db:   db,
		ctxm: client.NewPgxCtxManager(db),
	}
}
func (pm PgxManager) DefaultTrOrDb(ctx context.Context) (client.Client, error) {
	ctxClient := pm.ctxm.Default(ctx)
	if ctxClient != nil {
		return ctxClient, nil
	}
	return pm.db, nil
}

func (pm PgxManager) TrOrDb(ctx context.Context, key interface{}) (client.Client, error) {
	ctxClient := pm.ctxm.ByKey(ctx, key)
	if ctxClient != nil {
		return ctxClient, nil
	}
	return pm.db, nil
}
