package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shivammahatha/initializers"
	"github.com/shivammahatha/models"
)


type Posts struct {
    UserId int    `json:"userId"`
    Id int `json:"id"`
    Title string `json:"title"`
	Body string `json:"body"`
}

func CallMultiUrl(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	

//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   log.Printf(sb)

  
   var pst  [] Posts

// decoding JSON array to
    // the country array
    err1 := json.Unmarshal(body, &pst)

   if err1 != nil {
   // if error is not nil
        fmt.Println(err1)
    }

	// printing decoded array
    // values one by one
    for i := range pst {
		 c := models.Apipost{}
		 c.UserId = pst[i].UserId 
		 c.Title = pst[i].Title 
		 c.Body = pst[i].Body 
						
		err = initializers.DB.Create(&c).Error
									
		fmt.Println(err)			

    }



}
