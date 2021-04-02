package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

)

func main() {
	fmt.Println("start")
	r := gin.Default()
	r.GET("get", func(c *gin.Context) {
		c.String(200,"get");
	})

	r.POST("post", func(c *gin.Context) {
		c.String(200,"post");
	})

	r.Handle("DELETE", "/delete",func(c *gin.Context) {
		c.String(200,"delete");
	})
	r.Run()
}