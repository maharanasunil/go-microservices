package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunil/go-crud/handlers"
	"github.com/sunil/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	// 1. Setup Gin
	r := gin.Default()

	// 2. Create routes
	r.POST("/users", handlers.CreateUser)

	r.POST("/posts", handlers.PostsCreate)
	r.GET("/posts", handlers.PostsFetch)
	r.GET("/posts/:id", handlers.PostsShow)

	r.PUT("/posts/:id", handlers.PostsUpdate)

	r.DELETE("/posts/:id", handlers.PostsDelete)

	// 3. Start server
	r.Run(":8080")
}
