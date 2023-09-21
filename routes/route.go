package route

import (
	"github.com/faysalahmed-dev/go-microservice-order-api/handlers"
	"github.com/go-chi/chi/v5"
);

func LoadRoutes() *chi.Mux {
	r :=  chi.NewRouter()
	r.Get("/", handlers.GetAllOrders)
	r.Get("/{id}", handlers.GetOrder)
	r.Post("/", handlers.CreateOrder)
	r.Patch("/{id}", handlers.UpdateOrder)
	r.Delete("/{id}", handlers.DeleteOrder)
	return r
}
