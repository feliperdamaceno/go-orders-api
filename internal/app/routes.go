package app

import (
	"net/http"

	"github.com/feliperdamaceno/go-orders-api/internal/handler"
	"github.com/feliperdamaceno/go-orders-api/internal/repository/order"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *App) LoadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", func(router chi.Router) {
		order := &handler.OrderHandler{
			Repo: &order.RedisRepo{
				Client: a.Rds,
			},
		}

		router.Post("/", order.Create)
		router.Get("/", order.GetAll)
		router.Get("/{id}", order.GetById)
		router.Put("/{id}", order.UpdateById)
		router.Delete("/{id}", order.DeleteById)
	},
	)

	a.Router = router
}
