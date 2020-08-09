package main

import (
	"math"
	"strconv"
)

func euler26_2(quotient int) int {
	var exp = len(strconv.Itoa(quotient))
	remainder := int(math.Pow10(exp))
	var maxRecurrLength = 0
	var maxRecurrLengthQuotient = 0
	for i := quotient; i > quotient-remainder/10; i-- {
		var remainders = make(map[int]int)
		var r = remainder
		for pos := 1; ; pos++ {
			r = 10 * (r % i)
			v, ok := remainders[r]
			if ok {
				if pos-v > maxRecurrLength {
					maxRecurrLength = pos - v
					maxRecurrLengthQuotient = i
				}
				break
			}
			remainders[r] = pos
		}
	}
	return maxRecurrLengthQuotient
}
