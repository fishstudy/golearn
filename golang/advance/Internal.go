package main

import "fmt"

func main() {
	getLen()
	closeChan()
}

func getLen() {
	mIdSliceDes := make([]string, 2, 9)
	mIdSliceDes[0] = "id-des-1"
	mIdSliceDes[1] = "id-des-2"
	//mIdSliceDes[2] = "id-des-3"
	fmt.Println("mIdSliceDes len:", len(mIdSliceDes))
	fmt.Println("mIdSliceDes cap:", cap(mIdSliceDes))

}

func closeChan() {
	mChan := make(chan int, 1)
	defer close(mChan)
	mChan <- 1

}
