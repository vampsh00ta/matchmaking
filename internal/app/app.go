package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"matchmaking/config"
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
	redrep := redisorep.New(clientRedis)
	//r, err := psqlrep.GetRating(ctx, 1)
	//fmt.Println(r, err)
	//rep = rep
	//
	//matchmaking, err := rep.GetRating(ctx, 100)
	//fmt.Println(matchmaking, err)
	//
	srvc := service.New(psqlrep)
	//
	//trptLogger := logger.Sugar()
	//handlers := httphandler.New(srvc, trptLogger)
	//
	//logger.Info("Listening...")
	//server := &http.Server{Addr: ":" + cfg.HTTP.Port, Handler: handlers, ReadHeaderTimeout: 2 * time.Second}
	//
	//go func() {
	//	if err := server.ListenAndServe(); err != nil {
	//		logger.Fatal(err.Error())
	//	}
	//}()
	//stop := make(chan os.Signal, 1)
	//signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
	//<-stop
}
