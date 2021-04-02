package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("");
	r := gin.Default()
	r.POST("/test",func(c *gin.Context){
		bodyByts , err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		c.String(http.StatusOK,string(bodyByts))
	})
	r.Run()
}