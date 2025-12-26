package problems

import "fmt"

func max(dict map[string]int) string {
	max := 0
	var maxKey string
	for key, value := range dict {
		if value > max {
			max = value
			maxKey = key
		}
	}
	return maxKey
}

func Max_value() {
	dict := map[string]int{"a": 11, "b": 2, "c": 3}
	fmt.Println(max(dict))
}