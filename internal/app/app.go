package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/feliperdamaceno/go-orders-api/config"
	"github.com/feliperdamaceno/go-orders-api/internal/handler"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rds    *redis.Client
}

func New() *App {
	config.LoadConfig()

	app := &App{
		router: handler.GetRoutes(),
		rds:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	port := config.Config.PORT
	host := config.Config.HOST

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: a.router,
	}

	err := a.rds.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to initialize redis: %w", err)
	}

	fmt.Printf("server running on http://%v:%v", host, port)
	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to initialize server: %w", err)
	}

	return nil
}
