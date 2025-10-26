package application

import (
	"context"
	"fmt"
	"net/http"
)

// Defines an App type that stores the router
type App struct {
	router http.Handler
}

// initializes App and sets router by calling loadRoutes()
func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
