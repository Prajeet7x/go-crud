package controllers

import (
	"crud/crud-module/initializers"
	model "crud/crud-module/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Create a post
	post := model.Post{Title: "Hey", Body: "First post"}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
