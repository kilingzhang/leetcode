package main

import "fmt"

func validPalindrome(s string) bool {

	l := 0
	r := len(s) - 1
	skip := 0

	flag := true
	for l < r {
		if skip != 0 && s[l] != s[r] {
			flag = false
		}

		if s[l] != s[r] {
			skip++
			r++
		}

		l++
		r--
	}

	if flag {
		return true
	}

	l = 0
	r = len(s) - 1
	skip = 0

	flag = true
	for l < r {
		if skip != 0 && s[l] != s[r] {
			flag = false
		}

		if s[l] != s[r] {
			skip++
			l--
		}

		l++
		r--
	}

	if flag {
		return true
	}

	return false
}

func main() {

	fmt.Println(validPalindrome("cbbcc"))

}
