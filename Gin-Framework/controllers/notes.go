package controllers

import "github.com/gin-gonic/gin"

type NotesController struct{}

// We create a function which takes a router as input and can access all struct methods of NotesController Struct
func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")

	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": "Get request for notes",
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": "Post request for notes",
		})
	}
}
