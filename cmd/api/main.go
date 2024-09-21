package main

import (
	"context"
	"log"

	"github.com/feliperdamaceno/go-orders-api/internal/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
