package problems

import (
	"fmt"
)

func ReverseString() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scan(&input)
	r := []rune(input)
    i, j := 0, len(r)-1

    for i < j {
        r[i], r[j] = r[j], r[i]
        i++
        j--
    }

    fmt.Println("Reversed string:", string(r))
}