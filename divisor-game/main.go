package main

import tool "github.com/kilingzhang/leetcode"

func divisorGame(N int) bool {
	return N%2 == 0
}

func main() {
	tool.Dump(divisorGame(2))
}
