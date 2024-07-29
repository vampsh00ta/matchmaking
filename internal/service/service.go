package service

import (
	ipsql "matchmaking/internal/app/repository/psql"
	iredis "matchmaking/internal/app/repository/redis/v2"

	isrvc "matchmaking/internal/app/service"
)

type service struct {
	psql  ipsql.Repository
	redis iredis.Repository
}

func New(psql ipsql.Repository, redis iredis.Repository) isrvc.Matchmaking {
	return &service{
		psql:  psql,
		redis: redis,
	}
}
