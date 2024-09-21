package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feliperdamaceno/go-orders-api/config"
	"github.com/feliperdamaceno/go-orders-api/internal/handler"
)

func main() {
	config.LoadConfig()

	port := config.Config.PORT
	host := config.Config.HOST

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", host, port),
		Handler: handler.GetRoutes(),
	}

	fmt.Printf("Server running on http://%v:%v", host, port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
