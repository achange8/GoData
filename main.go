package main

import (
	"fmt"

	"github.com/achange8/Godata/sortdata"
)

func main() {
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}
	fmt.Println("arr =", arr)
	fmt.Println("sorted1 = ", sortdata.CountSort1(arr))
	fmt.Println("sorted2 = ", sortdata.CountSort2(arr))
}
