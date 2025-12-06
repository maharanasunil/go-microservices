package main

import (
	"github.com/sunil/go-crud/initializers"
	"github.com/sunil/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// This will create a table in our database
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Post{})
}
