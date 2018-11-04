package main

import (
	"fmt"

	"github.com/boris.voychev/leet.code/go/easy/valid.parentheses/stack"
)

func main() {
	fmt.Println(isValid("[]"))
}

func isValid(s string) bool {
	st := stack.New()
	dict := make(map[rune]rune)
	dict[')'] = '('
	dict['}'] = '{'
	dict[']'] = '['
	if len(s) == 1 {
		return false
	}
	for _, v := range s {
		if val, ok := dict[v]; ok {
			var topElement interface{}
			if st.IsEmpty() {
				topElement = '#'
			} else {
				topElement = st.Pop()
			}

			if topElement != val {
				return false
			}
		} else {
			st.Push(v)
		}
	}
	return st.IsEmpty()
}
