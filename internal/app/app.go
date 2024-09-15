package app

import (
	"context"
	redisrep "matchmaking/internal/repository/redis"
	"net"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"matchmaking/config"

	grpchandlers "matchmaking/internal/handler/grpc"
	"matchmaking/internal/service"

	//httphandler "matchmaking/internal/handler/http"
	psqlrep "matchmaking/internal/repository/postgres_repository"
	"matchmaking/pkg/client"
)

func NewLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logger.Sync()
	}()
	sugar := logger.Sugar()
	return sugar
}
func Run(cfg *config.Config) {
	ctx := context.Background()
	defer ctx.Done()

	sugar := NewLogger()
	pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		sugar.Fatalf("matchmaking - Run - postgres.New: %v", err)
	}
	defer pg.Close()
	ratingRep := psqlrep.NewRating(pg)
	txManager := psqlrep.NewPgxTxManager(pg)

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	queueRep := redisrep.NewQueue(clientRedis)

	matchSrvc := service.NewMatch(ratingRep, txManager, queueRep)

	lis, err := net.Listen("tcp", ":"+cfg.HTTP.Port)
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	matchMakingServer := grpchandlers.New(matchSrvc, sugar)
	grpchandlers.Register(s, matchMakingServer)
	sugar.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}
}
