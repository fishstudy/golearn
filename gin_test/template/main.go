package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Booking struct {
	CheckIn time.Time `form:"check_in" binding:"required" time_format:"2006--1-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-11-02"`
}

func main() {
	r := gin.Default()
	r.GET("/bookable",testing)
	r.POST("/test",testing)
	r.Run()
}

func testing(c *gin.Context)  {
	var book Booking
	if err := c.ShouldBind(&book); err==nil {
		c.JSON(200,gin.H{"message":"OK"})
	} else {
		c.JSON(500,gin.H{"error":err.Error()})
		c.Abort()
		return
	}

}