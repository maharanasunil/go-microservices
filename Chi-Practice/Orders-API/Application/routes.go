package application

import (
	"net/http"

	handler "github.com/chi/orders-api/Handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func loadRoutes() *chi.Mux {
	// create a new instance of chi router
	router := chi.NewRouter()
	// logging each request
	router.Use(middleware.Logger)

	// We write an anonymous function
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	// Setup a sub router. All routes inside loadOrderRoutes are prefixed with /orders
	router.Route("/orders", loadOrderRoutes)

	return router
}

// loadOrderRoutes takes a chi.Router and registers 5 RESTful endpoints for orders
func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)

}
