package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("Useage calc command ..")
	fmt.Println("\n the commands ar ")
}

func main() {
	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("useage calc add <inter1> <inter2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("useage calc add <inter1> <inter2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result :", ret)
	case "sqort":
		if len(args) != 2 {
			fmt.Println("useage sqrt <inter>")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("useage calc sqort inter")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result :", ret)
	default:
		Usage()
	}
}
