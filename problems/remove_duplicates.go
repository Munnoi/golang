package problems

import (
	"fmt"
)

func RemoveDuplicates() {
	arr := []int{1, 2, 2, 3, 4, 4, 5}

	seen := make(map[int]struct{}) // 'seen' is a map of key 'int' and value 'struct{}'.
	result := []int{}

	for _, v := range arr {
		if _, ok := seen[v]; !ok { // 'ok' becomes false if the value (v) is not found in the seen map.
			seen[v] = struct{}{} // Assigning an empty struct{} to the value in map.
			result = append(result, v) // Appending the unique value to the results array.
		}
	}
	fmt.Println(result)
}