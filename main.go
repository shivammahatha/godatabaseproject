package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shivammahatha/controller"
	"github.com/shivammahatha/initializers"
)
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDataBase()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controller.Postscreate)
	r.GET("/posts", controller.PostsIndex)
	r.GET("/posts/:id", controller.PostsShow)
	r.PUT("/posts/:id", controller.PostsUpdate)
	r.DELETE("/posts/:id", controller.PostsDelete)

	r.Run()
}