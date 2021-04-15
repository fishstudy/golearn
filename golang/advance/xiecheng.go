package main

import (
	"fmt"
	"learngo/gorutime"
	"learngo/syncwg"
	"runtime"
	"time"
)

func Loop() {
	fmt.Printf("cpu num=%d", runtime.NumCPU())
	fmt.Println("")
	runtime.GOMAXPROCS(runtime.NumCPU())
	go gorutime.Loop()
	go gorutime.Loop()
}
func channel() {
	go gorutime.Send()
	go gorutime.Reciver()
}
func syncWg() {
	syncwg.Read()
	go syncwg.Write()
	syncwg.Wg.Wait()
	fmt.Println("all done")
}
func main() {
	syncWg()

	time.Sleep(time.Second * 9)
}
