package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	res := euler26_2(15)
	fmt.Println(res)

	duration := time.Since(start)
	fmt.Println(duration)
}
