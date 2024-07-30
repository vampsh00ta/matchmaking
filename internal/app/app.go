package app

import (
	"context"
	"net"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"matchmaking/config"

	grpchandlers "matchmaking/internal/handler/grpc"
	redisrep "matchmaking/internal/repository/redis/v2"
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
	psqlrep := psqlrep.New(pg)

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	redrep := redisrep.New(clientRedis)

	srvc := service.New(psqlrep, redrep)

	lis, err := net.Listen("tcp", ":"+cfg.HTTP.Port)
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	matchMakingServer := grpchandlers.New(srvc, sugar)
	grpchandlers.Register(s, matchMakingServer)
	sugar.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}
}
