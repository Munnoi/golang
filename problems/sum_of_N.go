package problems

import (
	"fmt"
)

func v1(N int) int {
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	return sum
}

func v2(N int) int {
	return N * (N + 1) / 2
}

func Sum_of_N() {
	var N int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&N)

	sum := v2(N)

	fmt.Println("Sum = ", sum)
}