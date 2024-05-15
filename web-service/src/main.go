package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	 "web-service/random-data-module"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"version": "1.0.0",
		})
	})
	fmt.Println("Server is running on port 8080")
	r.Run() // listen and serve on
}