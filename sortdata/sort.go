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
	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return sorted
}

func MapSort() {
	students := []struct {
		Name   string
		Height float64
	}{
		{Name: "Kyle", Height: 173.4},
		{Name: "Ken", Height: 164.5},
		{Name: "Ryu", Height: 178.8},
		{Name: "Honda", Height: 154.2},
		{Name: "Hwarang", Height: 188.8},
		{Name: "Lebron", Height: 209.8},
		{Name: "Hodong", Height: 197.7},
		{Name: "Tom", Height: 164.8},
		{Name: "Kevin", Height: 164.8},
	}

	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	//140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name:", name, "height:", float64(i)/10)
		}
	}
}
