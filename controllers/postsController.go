package controllers

import (
	"crud/crud-module/initializers"
	model "crud/crud-module/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := model.Post{Title: body.Title, Body: body.Body}

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

func GetAllPosts(c *gin.Context) {

	var posts []model.Post

	//Get the posts
	initializers.DB.Find(&posts)

	//Return the data
	c.JSON(200, gin.H{
		"All posts": posts,
	})
}
