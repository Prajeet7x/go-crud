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

func GetOnePost(c *gin.Context) {
	//Get id of URL
	id := c.Param("id")

	// Get the posts
	var post []model.Post

	// Get the first record ordered by primary key
	initializers.DB.First(&post, id)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	c.JSON(200, gin.H{
		"All posts": post,
	})
}

func DeletePost(c *gin.Context) {
	//Get the id off form the URL
	id := c.Param("id")

	var post []model.Post

	initializers.DB.Delete(&post, id)
	// DELETE from emails where id = 10;

	c.JSON(200, gin.H{
		"Deleted post": post,
	})
}
