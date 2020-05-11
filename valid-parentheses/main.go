package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func isValid(s string) bool {

	if s == "" {
		return true
	}

	stack := []uint8{s[0]}
	for i := 1; i < len(s); i++ {
		if len(stack) != 0 && stack[len(stack)-1] == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
			continue
		}
		if len(stack) != 0 && stack[len(stack)-1] == '[' && s[i] == ']' {
			stack = stack[:len(stack)-1]
			continue
		}
		if len(stack) != 0 && stack[len(stack)-1] == '{' && s[i] == '}' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}

func main() {
	tool.Dump(isValid("()[]{}"))
}
