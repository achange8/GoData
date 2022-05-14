package main

import datastrucks "main/dataStrucks"

func main() {
	list := &datastrucks.LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.PrintNode()
	list.DeleteNode(list.Root.Next)
	list.PrintNode()
	list.DeleteNode(list.Tail)
	list.PrintNode()

}
