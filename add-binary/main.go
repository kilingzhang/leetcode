package main

import (
	"fmt"
	tool "github.com/kilingzhang/leetcode"
)

func addBinary(a string, b string) string {
	return ""
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
	tool.Dump(addBinary("11", "1"))
}
