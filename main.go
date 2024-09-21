package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", rootHandler)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server running on http://localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server workng as expected!"))
}
