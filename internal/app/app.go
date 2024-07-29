package app

import (
	"context"
	"fmt"
	"log"
	"matchmaking/internal/handler/grpc/pb"
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

func Run(cfg *config.Config) {
	ctx := context.Background()
	defer ctx.Done()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = logger.Sync()
	}()
	pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		logger.Fatal(fmt.Sprintf("clean - Run - postgres.New: %w", err.Error()))
	}
	defer pg.Close()
	psqlrep := psqlrep.New(pg)

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
	redrep := redisrep.New(clientRedis)
	//err = redrep.AddUserToQueue(ctx, entity.User{1, 1})
	//fmt.Println(err)

	//u, err := redrep.GetUsersInQueue(ctx)
	//fmt.Println(u, err)
	//r, err := psqlrep.GetRating(ctx, 1)
	//fmt.Println(r, err)
	//rep = rep
	//
	//matchmaking, err := rep.GetRating(ctx, 100)
	//fmt.Println(matchmaking, err)
	//
	srvc := service.New(psqlrep, redrep)
	//fmt.Println(srvc.MatchResult(ctx, 1, 2))
	//
	//trptLogger := logger.Sugar()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	matchMakingServer := grpchandlers.New(srvc)
	pb.RegisterMatchmakingServer(s, matchMakingServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
