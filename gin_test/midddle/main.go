package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io"
)

func main() {
	fmt.Println("");
	f,_:= os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/test", func(c *gin.Context) {
		name := c.DefaultQuery("name","default_name")
		panic("test panic")
		c.String(200,"%s",name)
	})
	r.Run()
}