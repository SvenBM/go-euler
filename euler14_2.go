package main

func euler14_2(limit int) int {
	var maxCount, startingNum int = 0, 0
	var cache = make([]int, limit)

	for i := 2; i <= limit; i++ {
		count := 0
		for seq := i; seq > 1; {
			if seq < i {
				count += cache[seq-1]
				break
			}
			count++
			if seq%2 == 0 {
				seq /= 2
			} else {
				seq = 3*seq + 1
			}
		}
		cache[i-1] = count
		if count > maxCount {
			maxCount = count
			startingNum = i
		}
	}
	return startingNum
}
