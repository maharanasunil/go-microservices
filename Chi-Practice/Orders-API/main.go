package main

import (
	"context"
	"fmt"

	application "github.com/chi/orders-api/Application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
