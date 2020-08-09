package main

import (
	"math"
)

func euler12(divisorCount int) int {
	trianNum := 1
	for i := 2; ; i++ {
		trianNum += i
		divisors := 2
		sqrt := int(math.Sqrt(float64(trianNum)))
		for d := 2; d < sqrt; d++ {
			if trianNum%d != 0 {
				continue
			}
			divisors += 2
		}
		if sqrt+sqrt == trianNum {
			divisors++
		}

		//fmt.Printf("Trianguläre Zahl %v hat %v Teiler\n", trianNum, divisors)

		if divisors > divisorCount {
			break
		}
	}
	return trianNum
}
