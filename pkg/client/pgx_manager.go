package client

import "context"

type DbManager interface {
	DefaultTrOrDb(ctx context.Context) (Client, error)
	TrOrDb(ctx context.Context, key interface{}) (Client, error)
}

type TxFunc func(ctx context.Context) error
type Manager interface {
	Do(ctx context.Context, f TxFunc) error
}
