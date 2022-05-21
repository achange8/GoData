package main

import (
	"fmt"
	"regexp"
)

func main() {

	validEmail, _ := regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)
	fmt.Println(validEmail.MatchString("gildong@example.com"))

	// stack2 := dataStrucks.NewStack()
	// for i := 1; i <= 5; i++ {
	// 	stack2.Push(i)
	// }
	// fmt.Println("New Stack")
	// for !stack2.Empty() {
	// 	val := stack2.Pop()
	// 	fmt.Printf("%d ->", val)
	// }
}
