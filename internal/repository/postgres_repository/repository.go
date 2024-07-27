package postgresrep

import (
	ipsql "matchmaking/internal/app/repository/psql"
	"matchmaking/pkg/client"
)

type pg struct {
	client client.Client
}

func New(client client.Client) ipsql.Repository {
	return &pg{
		client,
	}
}
