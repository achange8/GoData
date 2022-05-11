package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node
	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}
	PrintNode(root)

	root, tail = DeleteNode(root, root, tail)
	PrintNode(root)

}

func AddNode(tail *Node, val int) *Node {

	node := &Node{val: val}
	tail.next = node
	return node
}

func DeleteNode(del *Node, root *Node, tail *Node) (*Node, *Node) {
	if del == root {
		root = root.next
		if root == nil {
			tail = nil
		}
		return root, tail
	}

	prev := root
	for prev.next != del {
		prev = prev.next
	}
	if del == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}

	println("Done!")
	return root, tail

}

func PrintNode(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d >", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
