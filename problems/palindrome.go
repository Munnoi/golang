package problems

import (
	"fmt"
)

func Palindrome() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	r := []rune(input)
	i, j := 0, len(r)-1

	for i < j {
		if r[i] != r[j] {
			fmt.Println("Not a palindrome")
			return
		}
		i++
		j--
	}
	fmt.Println("Palindrome")
}