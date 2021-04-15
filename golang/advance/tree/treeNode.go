package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	var node treeNode
	fmt.Println(node)
	node = treeNode{value: 3}
	node.left = &treeNode{}
	node.right = &treeNode{5, nil, nil}
}
