package main

func euler14(limit int) int {
	var maxCount, startingNum int = 0, 0
	for i := limit / 2; i <= limit; i++ {
		count := collatz(i, 0)
		if count > maxCount {
			maxCount = count
			startingNum = i
		}
	}
	return startingNum
}

func collatz(num, count int) int {
	if num == 1 {
		return count
	}
	count++
	if num%2 == 0 {
		return collatz(num/2, count)
	}
	return collatz(3*num+1, count)
}
