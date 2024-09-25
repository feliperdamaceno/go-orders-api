package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/feliperdamaceno/go-orders-api/internal/app"
)

func main() {
	app := app.New()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
