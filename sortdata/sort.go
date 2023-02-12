package sortdata

import "fmt"

func CountSort1(arr []int) []int {
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	fmt.Println("count : ", count)
	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}

	return sorted
}

func CountSort2(arr []int) []int {
	var count [11]int
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}
	fmt.Println("count : \n", count)
	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return sorted
}
