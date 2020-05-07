package main

import (
	tool "github.com/kilingzhang/leetcode"
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	max := 1

	for i := 1; i < len(s); i++ {
		j := i
		t := 1

		if len(s)-i < max {
			break
		}

		for {
			if strings.Contains(s[i-1:j], s[j:j+1]) {
				if t > max {
					max = t
				}
				break
			}
			j++
			t++
			if j >= len(s) {
				if t > max {
					max = t
				}
				break
			}
		}
	}
	return max
}

func main() {
	tool.Dump(lengthOfLongestSubstring("bbbbb"))
}
