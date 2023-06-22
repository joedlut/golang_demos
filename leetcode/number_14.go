package main

import "math"

func longestCommonPrefix(strs []string) string {
	min, index := MinLengthAndIndex(strs)
	m := make(map[byte]int)
	Pos := -1
	for i := 0; i < min; i++ {
		for j := 0; j <= len(strs)-1; j++ {
			m[strs[j][i]]++
		}
		if m[strs[index][i]] == len(strs) {
			Pos++
			delete(m, strs[index][i])
		} else {
			break
		}
	}
	if Pos == -1 {
		return ""
	} else {
		return strs[index][:Pos+1]
	}
}

func MinLengthAndIndex(strs []string) (m int, n int) {
	min := math.MaxInt32
	index := 0
	for currentIndex, str := range strs {
		l := len(str)
		if min > l {
			min = l
			index = currentIndex
		}
	}
	return min, index
}
