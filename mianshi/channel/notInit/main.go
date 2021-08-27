package main

func main() {
	var ch chan int
	ch <- 1

	//var c chan int
	//num, ok := <-c
	//fmt.Printf("读chan的协程结束, num=%v, ok=%v\n", num, ok)
}
