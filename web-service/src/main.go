package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"randomDataModule"
	"net/http"
	"modelModule"
)





func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"version": "1.0.0",
		})
	})
	r.GET("/random-str/:type", GetRandomAdditonOrSubtractionHandler)
	fmt.Printf("Server is running on port %d\n", 8089)
	randomDataModule.GetRandomDomain()
	r.Run(":8089") // listen and serve on
}


func createResponse(c *gin.Context,success bool, message string, data interface{}) {
	c.JSON(http.StatusOK, modelModule.Response{
	   Success: success,
	   Message: message,
	   Data: data,
   })
}



// 获取随机的加减法题目
func GetRandomAdditonOrSubtractionHandler(c *gin.Context) {
	rand_type := c.Param("type")

	switch rand_type {
		case "add-or-sub":
			createResponse(c,true, "success", randomDataModule.GetRandomAdditonOrSubtraction())
		default:
			createResponse(c,false, "type not found", nil)
	}
	
}