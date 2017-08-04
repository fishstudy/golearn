package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.ReadRequest()){
	r.ParseFrom()
	fmt.Println(r.From)
	fmt.Println("jpath", r.URL.Path)
}