package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	farray := []string{"flower", "flow", "flight"}
	sarray := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(farray))
	fmt.Println(longestCommonPrefix(sarray))
	fmt.Println(longestCommonPrefix([]string{}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minl := math.MaxInt16
	mins := ""
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < minl {
			minl = len(strs[i])
			mins = strs[i]
		}
	}

	result := ""
	for i := 0; i <= minl; i++ {
		c := mins[:i]
		areEqual := true
		for j := 0; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], c) {
				areEqual = false
			}
		}
		if areEqual {
			result = c
		}
	}

	return result
}
