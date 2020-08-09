package main

import (
	"math"
)

func euler10_2(number int64) int64 {
	candidates := make([]bool, number/2)
	var sum int64 = 2
	var sqrt int64 = int64(math.Sqrt(float64(number)))
	for i := int64(1); i < sqrt/2; i++ {
		if candidates[i] {
			continue
		}
		prime := 2*i + 1
		sum += prime
		for j := int64(i + prime); j < number/2; j += prime {
			candidates[j] = true
		}
	}

	for i := int64(sqrt / 2); i < number/2; i++ {
		if candidates[i] {
			continue
		}
		sum += 2*i + 1
	}
	return sum
}
