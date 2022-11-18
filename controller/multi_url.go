package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallMultiUrl(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	// var data struct {
	// 	userId string
	// 	title string
	// 	body string
		
	// }
	// c.Bind(&data)
	// //create a post
	// description := models.Description{data.userId: data.userId, title: data.title, body: data.body}
	//  result := initializers.DB.Create(&description)


	//  if result.Error != nil {
	// 	c.Status(400)
	// 	return
	//  }
	
   if err != nil {
      log.Fatalln(err)
   }
//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   log.Printf(sb)
}