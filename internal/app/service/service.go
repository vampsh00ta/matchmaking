package service

import "context"

type Matchmaking interface {
	FindMatch(ctx context.Context, tgID int) (int, error)
	MatchResult(ctx context.Context, tgIDWinner, tgIDLoser int) error
}
