package main

import (
	"github.com/gin-gonic/gin")

type Person struct {
	Name string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
	Age int `form:"age" binding:"required,gt=10"`
}

func main() {
	r := gin.Default()
	r.GET("/test",testing)
	r.POST("/test",testing)
	r.Run()
}

func testing(c *gin.Context)  {
	var person Person
	if err := c.ShouldBind(&person); err==nil {
		c.String(200,"%v",person)
	} else {
		c.String(500,"error %v",err)
		c.Abort()
		return
	}

}