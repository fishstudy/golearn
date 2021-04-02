package main


import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("")
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/static",http.Dir("static"))
	r.StaticFile("/favicon.ico","./favicon.ico")
	r.Run()
}