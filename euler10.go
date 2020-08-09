package main

func euler10(number int64) int64 {
	candidates := make([]bool, number/2)
	var sum int64 = 2
	for i := int64(1); i < number/2; i++ {
		if candidates[i] {
			continue
		}
		prime := 2*i + 1
		sum += prime
		for j := int64(i + prime); j < number/2; j += prime {
			candidates[j] = true
		}
	}
	return sum
}
