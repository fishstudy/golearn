package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("yuxuefeng", "feng")
	fmt.Println("yuxuefeng:", os.Getenv("yuxuefeng"))
	fmt.Println("sunyanyan:", os.Getenv("sunyanyan"))
	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "=>", pair[1])
	}
}
