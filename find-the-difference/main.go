package main

import (
	tool "github.com/kilingzhang/leetcode"
)

func findTheDifference(s string, t string) (ans byte) {
	l := len(s)
	ans = t[l]
	for i := 0; i < l; i++ {
		ans ^= s[i] ^ t[i]
	}
	return ans
}

func main() {
	tool.Dump(findTheDifference("abcd", "abcde"))
}
