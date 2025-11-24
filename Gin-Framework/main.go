package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunil/gin-practice/controllers"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint. It takes the path and gin.Context
	// router.GET("/ping", getJSONdata)
	// router.GET("/me/:id/:newId", getId)
	// router.POST("/me", getEmailPswd)

	// // PUT is basically used to update entire data
	// router.PUT("/upload", getEmailPswd)
	// // PATCH is used to update a small part
	// router.PATCH("/me", getEmailPswd)

	// // DELETE is basically used to delete a key/id
	// router.DELETE("/delete/:id", deleteId)

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":8000")
}

func getJSONdata(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message":      "pong from gin via Sunil",
		"new_variable": "This is a json data",
	})
}

func getId(c *gin.Context) {
	// Read path parameters
	var id = c.Param("id")
	var newId = c.Param("newId")

	c.JSON(200, gin.H{
		"user_id": id,
		"New_Id":  newId,
	})
}

func getEmailPswd(c *gin.Context) {
	type MeRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password"`
	}

	var meReq MeRequest
	// BindJSON - To serialize body json with struct
	if err := c.BindJSON(&meReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"email":    meReq.Email,
		"password": meReq.Password,
	})

}

func deleteId(c *gin.Context) {
	var id = c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "deleted!",
	})
}
