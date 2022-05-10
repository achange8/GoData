package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node

	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		AddNode(root, i)
	}
	node := root

	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func AddNode(root *Node, val int) {
	var endP *Node
	endP = root
	for endP.next != nil {
		endP = endP.next
	}

	node := &Node{val: val}
	endP.next = node
}
