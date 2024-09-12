package service

import "context"

type Match interface {
	Find(ctx context.Context, tgID int) (int, error)
	Result(ctx context.Context, tgIDWinner, tgIDLoser int) error
}
