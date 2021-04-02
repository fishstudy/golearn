package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
)


func main() {
	r := gin.Default()
	r.GET("/test",testing)
	srv := &http.Server{
		Addr : ":8085",
		Handler:r,
	}
	go func() {
		if err:=srv.ListenAndServe();err!=nil && err!=http.ErrServerClosed{
			log.Fatalf("listen:%s\n",err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx);err!=nil{
		log.Fatal("server shutdown :",err)
	}
	log.Println("server exiting")
	r.Run()
}

func testing(c *gin.Context)  {
	time.Sleep(10*time.Second)
	c.String(200,"hello test")
}