package main

import "fmt"

/*
28. 找出字符串中第一个匹配项的下标
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。
*/

func subStr(haystack string, needle string) int {
	m, n := len(needle), len(haystack)
	next := make([]int, m)
	//求解next 数据，直接背过模板
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	for i, j := 0, 0; i < n; i++ {
		//循环
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}
		//匹配成功
		if j == m {
			return i - m + 1
		}
	}
	fmt.Println(next)
	return -1
}

// a a b b a a f
// 0 1 0 0 1 2 0
func main() {
	subStr("aabbacaaabbaafcde", "aabbaaf")

}
