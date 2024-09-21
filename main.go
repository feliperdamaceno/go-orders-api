package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: http.HandlerFunc(rootHandler),
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
