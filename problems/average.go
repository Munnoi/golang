package problems

import (
	"fmt"
	"math"
)

func RoundTo(n float64, decimals uint32) float64 {
	// Multiply the number by 10^decimals, round it, then divide back
	powerOf10 := math.Pow(10, float64(decimals))
	return math.Round(n*powerOf10) / powerOf10
}

func Avg() {
	arr := []float64{4, 1, 9, -2, 7, 3}

	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	avg := sum / float64(len(arr))

	count := 0
	for _, v := range arr {
		if v > avg {
			count++
		}
	}
	fmt.Println("Avg = ", RoundTo(avg, 2))
	fmt.Println("Count = ", count)
}