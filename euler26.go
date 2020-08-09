package main

import (
	"math"
	"strconv"
	"strings"
)

func euler26(quotient int) int {
	var exp = len(strconv.Itoa(quotient))
	remainder := int(math.Pow10(exp))
	var start = "0."
	for i := 1; i < exp; i++ {
		start += "0"
	}
	var maxRecurrLength = 0
	var maxRecurrLengthQuotient = 0
	for i := quotient; i > quotient-remainder/10; i-- {
		var rl = getRecurringLength(i, exp, remainder, start)
		if rl > maxRecurrLength {
			maxRecurrLength = rl
			maxRecurrLengthQuotient = i
		}
	}
	return maxRecurrLengthQuotient
}

func getRecurringLength(quotient, exp, remainder int, start string) int {
	res := start
	for {
		var div = remainder / quotient
		res += strconv.Itoa(div)
		remainder = 10 * (remainder % quotient)
		if remainder == 0 {
			// fmt.Println("exakt teilbar. ergebnis: " + res)
			return 0
		}
		var lastExpDigitsPos = len(res) - exp
		var lastExpDigits = res[lastExpDigitsPos:]
		var indexLastSeen = strings.Index(res[:len(res)-1], lastExpDigits)
		if indexLastSeen != -1 {
			// fmt.Println(res)
			return lastExpDigitsPos - indexLastSeen
		}
	}
}
