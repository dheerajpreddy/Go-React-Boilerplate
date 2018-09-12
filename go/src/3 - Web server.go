package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()                          // Initialize server with default settings
  r.GET("/", func(c *gin.Context) {           // Creates route for root(/), second argument is function to execute
    c.JSON(200, gin.H{
      "message": "Hello World",
    })
  })
  r.Run()                                     // By default, isten and serve on http://localhost:8080
}
