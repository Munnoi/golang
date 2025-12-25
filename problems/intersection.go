package problems

import (
	"fmt"
)

func Intersection() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 5, 6, 7}

	seen := make(map[int]bool) // `seen` is a map of key `int` and value `bool`.
	result := []int{} // An empty integer array.

	for _, v := range arr1 {
		seen[v] = true
	}

	for _, v := range arr2 {
		if seen[v] {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}