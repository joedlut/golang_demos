package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	length := len(xStr)
	for i := 0; i < length/2; i++ {
		if xStr[i] != xStr[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var a int
	a = -121
	fmt.Println(isPalindrome(a))
}
