package main

import (
	"strconv"
)

func euler24(index int) string {
	var faculties = map[int]int{1: 1}

	for i := 2; i <= 10; i++ {
		faculties[i] = i * faculties[i-1]
	}

	var pool = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var res = ""
	var remainder = index
	for stelle := 10; stelle > 1; stelle-- {
		var div = remainder / faculties[stelle-1]
		if div > stelle {
			res += strconv.Itoa(pool[0])
			pool = pool[1:]
			continue
		}
		res += strconv.Itoa(pool[div])
		pool = append(pool[:div], pool[div+1:]...)
		remainder = remainder % faculties[stelle-1]
		if remainder == 0 {
			for i := 0; i < len(pool); i++ {
				res += strconv.Itoa(pool[i])
			}
			break
		}
	}
	return res
}
