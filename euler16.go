package main

func euler16(exp int) int {
	curr := []byte{1}
	var uebertrag byte = 0
	for ; exp > 0; exp-- {
		for stelle := len(curr) - 1; stelle >= 0; stelle-- {
			sum := 2*curr[stelle] + uebertrag
			curr[stelle] = sum % 10
			uebertrag = sum / 10
		}
		if uebertrag > 0 {
			curr = append([]byte{uebertrag}, curr...)
			uebertrag = 0
		}
	}

	checksum := 0
	for _, s := range curr {
		checksum += int(s)
	}

	return checksum
}
