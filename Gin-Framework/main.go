package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint. It takes the path and gin.Context
	router.GET("/ping", getJSONdata)
	router.Run(":8000")
}

func getJSONdata(gin_context *gin.Context) {
	// Return JSON response
	gin_context.JSON(http.StatusOK, gin.H{
		"message":      "pong from gin via Sunil",
		"new_variable": "This is a json data",
	})
}
