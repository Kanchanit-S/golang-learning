package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", getData)
	r.GET("/testAir", getData)
	r.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello",
	})
}
