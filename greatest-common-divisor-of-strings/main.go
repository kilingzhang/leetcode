package main

import tool "github.com/kilingzhang/leetcode"

func main() {

	str1 := "ABABAB"
	str2 := "ABAB"
	tool.Dump(gcdOfStrings(str1, str2))

}

func gcdOfStrings(str1 string, str2 string) string {

	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(a int, b int) int {

	if a == b {
		return a
	}

	if a > b {
		return gcd(a-b, b)
	}

	return gcd(b-a, a)
}
