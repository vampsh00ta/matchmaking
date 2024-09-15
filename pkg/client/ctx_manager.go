package client

import "context"

type CtxTrKey struct{}

type CtxManager interface {
	SetDefault(ctx context.Context, tr Client) context.Context
	SetByKey(ctx context.Context, tr Client, key interface{}) context.Context
	Default(ctx context.Context) Client
	ByKey(ctx context.Context, key interface{}) Client
}

type PgxCtxManager struct {
	client Client
}

func NewPgxCtxManager(client Client) CtxManager {
	return &PgxCtxManager{client: client}
}
func (mngr PgxCtxManager) SetByKey(ctx context.Context, tr Client, key interface{}) context.Context {

	ctx = context.WithValue(ctx, key, tr)
	return ctx
}
func (mngr PgxCtxManager) SetDefault(ctx context.Context, tr Client) context.Context {

	ctx = context.WithValue(ctx, CtxTrKey{}, tr)
	return ctx
}

func (mngr PgxCtxManager) Default(ctx context.Context) Client {
	val := ctx.Value(CtxTrKey{})
	client, ok := val.(Client)
	if !ok {
		return nil
	}

	return client
}
func (mngr PgxCtxManager) ByKey(ctx context.Context, key interface{}) Client {
	val := ctx.Value(key)
	client, ok := val.(Client)
	if !ok {
		return nil
	}

	return client
}
