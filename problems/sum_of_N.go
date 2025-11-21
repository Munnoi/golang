package problems

import (
	"fmt"
)

func Sum_of_N() {
	var N int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&N)

	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}

	fmt.Println("Sum = ", sum)
}