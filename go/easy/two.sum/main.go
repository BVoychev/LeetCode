package main

import (
	"fmt"
)

func main() {
	givenNumbers := []int{1, 2, 4, 5}
	target := 5
	fmt.Println(twoSum(givenNumbers, target))
}

func twoSum(nums []int, target int) []int {
	type t struct {
		exists bool
		index  int
	}
	dict := make(map[int]t)
	dict[nums[0]] = t{
		exists: true,
		index:  0,
	}
	for i := 1; i < len(nums); i++ {
		complement := target - nums[i]
		if dict[complement].exists {
			return []int{dict[complement].index, i}
		}
		dict[nums[i]] = t{
			exists: true,
			index:  i,
		}
	}
	return []int{}
}
