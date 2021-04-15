package syncwg

import (
	"fmt"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func Read() {
	for i := 0; i < 3; i++ {
		Wg.Add(1)
	}
}
func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Done ->", i)
		Wg.Done()

	}
}
