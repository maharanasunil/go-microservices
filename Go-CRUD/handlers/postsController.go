package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunil/go-crud/initializers"
	"github.com/sunil/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get Data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a POST
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsFetch(c *gin.Context) {
	// Get the posts
	var posts []models.Post // Collect array of all the posts data
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id from path param
	id := c.Param("id")

	// Get Post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id from path param
	id := c.Param("id")

	// Get data from request body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find post which has to be updated
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id from path param
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Id: %s is deleted", id),
	})
}
