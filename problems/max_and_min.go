package problems

import (
	"fmt"
)

func Max_and_min() {
	arr := []int{4, 1, 9, -2, 7, 3}

	min, max := arr[0], arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println("Max = ", max)
	fmt.Println("Min = ", min)
}