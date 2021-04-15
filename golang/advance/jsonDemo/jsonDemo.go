package jsonDemo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIp   string
	ServerPort int
}

func Serialize() {
	server := new(Server)
	server.ServerName = "demo-for-json"
	server.ServerIp = "127.0.0.1"
	server.ServerPort = 8080
	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("Marshal err:", err.Error())
	}
	fmt.Println("Marsha1", string(b))
}
