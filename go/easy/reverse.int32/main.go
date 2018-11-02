package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(math.MaxUint32 + 1))
}

func reverse(x int) int {
	revNum := 0

	if x > 0 {
		for x > 0 {
			revNum = revNum*10 + x%10
			x = x / 10
		}
		if revNum > math.MaxInt32 {
			return 0
		}
		return revNum
	}
	for x < 0 {
		revNum = revNum*10 + x%10
		x = x / 10
	}
	if revNum < -math.MaxInt32 {
		return 0
	}
	return revNum
}
