package psql

import "context"

type Rating interface {
	GetByTgID(ctx context.Context, tgID int) (int, error)
	UpdateByTgID(ctx context.Context, tgID int, matchmakingChange int) error
	Init(ctx context.Context, tgID, matchmaking int) error
}
