package generics

import (
	"fmt"
)

func Sum[T int | float32 | float64](nums []T) T {
	var total T
	for _, v := range nums {
		total += v
	}

	return total
}

func Any[T any](slices []T) string {
	var result string
	i := 0
	for num, v := range slices {
		if i == num {
			result = fmt.Sprintf("%v", v)
		} else {
			result = fmt.Sprintf("%v, %v", result, v)
		}
	}

	return result
}
