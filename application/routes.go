package application

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yash91989201/go_microservice/handler"
)

func registerRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/order", registerOrderRoutes)

	return router
}

func registerOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Get("/", orderHandler.Create)
	router.Post("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetById)
	router.Put("/{id}", orderHandler.UpdateById)
	router.Delete("/{id}", orderHandler.Delete)
}
