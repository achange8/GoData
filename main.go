package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNode()
	list.DeleteNode(list.root.next)
	list.PrintNode()
	list.DeleteNode(list.tail)
	list.PrintNode()

}

func AddNode(tail *Node, val int) *Node {

	node := &Node{val: val}
	tail.next = node
	return node
}

func (l *LinkedList) DeleteNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
	node.next = nil
}

func (l *LinkedList) PrintNode() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}
