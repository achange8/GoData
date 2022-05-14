package main

func main() {
	list := &dataStructs.LinkedList{}
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
