package main



import (
	"github.com/gin-gonic/gin"
	"time"
)
type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2020-11-11"`
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
		c.String(200,"person bind error %v",err)
	}

}