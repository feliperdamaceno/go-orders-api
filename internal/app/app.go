package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/feliperdamaceno/go-orders-api/config"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Router http.Handler
	Rds    *redis.Client
}

func New() *App {
	app := &App{
		Rds: redis.NewClient(&redis.Options{}),
	}

	app.LoadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	config, err := config.NewConfig()

	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.Host, config.Port),
		Handler: a.Router,
	}

	err = a.Rds.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to initialize redis: %w", err)
	}

	defer func() {
		if err := a.Rds.Close(); err != nil {
			fmt.Println("failed to close redis: %w", err)
		}
	}()

	fmt.Printf("server running on http://%v:%v\n", config.Host, config.Port)

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to initialize server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, stop := context.WithTimeout(context.Background(), time.Second*10)
		defer stop()

		return server.Shutdown(timeout)
	}
}
