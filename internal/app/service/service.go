package service

import "context"

type Service interface {
	FindMatch(ctx context.Context, tgID int) (int, error)
	MatchResult(ctx context.Context, tgIDWinner, tgIDLoser int) error
}
