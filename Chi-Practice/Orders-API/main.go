package main

import (
	"context"
	"fmt"

	application "github.com/chi/orders-api/Application"
)

func main() {
	// constructs your application (creates the router and returns an *App)
	app := application.New()
	// starts the HTTP server on port :3000
	err := app.Start(context.TODO()) // context.TODO() is a placeholder context — used here just to satisfy the Start function signature
	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
