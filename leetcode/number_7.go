package main

import (
	"fmt"
	"math"
)

func sliceToNum(s []int) int {
	num := 0
	j := 1
	for i := len(s) - 1; i >= 0; i-- {
		num += j * s[i]
		j *= 10
	}
	return num
}

func reverse(x int) int {
	var flag bool = true
	if x < 0 {
		flag = false
		x = -x
	}
	s := make([]int, 0)
	for x > 0 {
		c := x % 10
		s = append(s, c)
		x = x / 10
	}
	result := sliceToNum(s)
	if result > math.MaxInt32 {
		return 0
	}
	if !flag {
		return -result
	}
	return result
}

func main() {
	fmt.Println(reverse(1534236469))
}
