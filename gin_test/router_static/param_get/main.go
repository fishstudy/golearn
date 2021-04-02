package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"bytes"
)

/*func main() {
	fmt.Println("");
	r := gin.Default()
	r.GET("test",func(c *gin.Context){
		firstName := c.Query("first_name")
		lastName := c.Query("last_name")
		c.String(http.StatusOK,"%s,%s",firstName,lastName)
	})
	r.Run()
}*/

func main()  {

	r := gin.Default()
	r.POST("/test5", func(c *gin.Context) {
		bodyByts,err:=ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest,err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		//c.String(http.StatusOK,string(bodyByts))
		firstName := c.PostForm("first_name")
		lastName  := c.PostForm("last_name")
		c.String(http.StatusOK,"%s,%s,%s",firstName,lastName,string(bodyByts))
	})
	r.Run()
}