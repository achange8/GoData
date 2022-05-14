package main

import (
	"fmt"
	"main/dataStrucks"
)

func main() {

	stack2 := dataStrucks.NewStack()
	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}
	fmt.Println("New Stack")
	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}
}
