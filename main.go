package main

import (
	"crud/crud-module/controllers"
	"crud/crud-module/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDataBase()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetOnePost)
	r.DELETE("posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:3000
}
