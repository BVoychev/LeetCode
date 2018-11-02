package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(123))
}

func isPalindrome(x int) bool {
	firstNum := x
	revNum := 0
	for x > 0 {
		revNum = revNum*10 + x%10
		x = x / 10
	}

	if firstNum == revNum {
		return true
	}
	return false
}
