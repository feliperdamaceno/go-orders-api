package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func LoadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", func(router chi.Router) {
		order := &OrderHandler{}

		router.Post("/", order.Create)
		router.Get("/", order.GetAll)
		router.Get("/{id}", order.GetById)
		router.Put("/{id}", order.UpdateById)
		router.Delete("/{id}", order.DeleteById)
	},
	)

	return router
}
