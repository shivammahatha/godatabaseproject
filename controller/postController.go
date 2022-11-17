package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shivammahatha/initializers"
	"github.com/shivammahatha/models"
)

func Postscreate(c *gin.Context) {
//get data off req body
var body struct {
	MobileNumber string
	Name string
}
c.Bind(&body)
//create a post
 post := models.Post{Name: body.Name, MobileNumber: body.MobileNumber}
 result := initializers.DB.Create(&post)

 if result.Error != nil {
	c.Status(400)
	return
 }

 //return it

	c.JSON(200, gin.H{
		"post": post,
	})
}
 func PostsIndex(c *gin.Context){
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them 
	c.JSON(200, gin.H{
		"posts": posts,
	})

 }
 func PostsShow(c *gin.Context){
	//get id off url
	id := c.Param("id")

	// Get the posts
	var post []models.Post
	initializers.DB.First(&post,id)

	//Respond with them 
	c.JSON(200, gin.H{
		"post": post,
	})

 }
 func PostsUpdate(c *gin.Context){
	//get id off url
	id := c.Param("id")

	// Get the data of request body
	var body struct {
		MobileNumber string
		Name string
	}
	c.Bind(&body)

	// find the post where updating
	var post []models.Post
	initializers.DB.First(&post,id)

	//updated it
	initializers.DB.Model(&post).Updates(models.Post{
		Name: body.Name,
		MobileNumber: body.MobileNumber,
	})
	//Respond with them 
	c.JSON(200, gin.H{
		"post": post,
	})

 }
 func PostsDelete(c *gin.Context){
	// get the id of url
	id := c.Param("id")

	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respond
	c.Status(200)
 }