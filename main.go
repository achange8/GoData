package main

import (
	"fmt"
	"main/dataStrucks"
)

func main() {
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	fmt.Println("queue")
	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}
	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

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
