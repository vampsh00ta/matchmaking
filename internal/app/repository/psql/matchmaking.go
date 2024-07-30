package psql

import "context"

type Matchmaking interface {
	GetRating(ctx context.Context, tgID int) (int, error)
	UpdateRating(ctx context.Context, tgID int, matchmakingChange int) error
	InitRating(ctx context.Context, tgID, matchmaking int) error
}
