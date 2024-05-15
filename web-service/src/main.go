package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"randomDataModule"
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
	randomDataModule.GetRandomDomain()
	r.Run() // listen and serve on
}