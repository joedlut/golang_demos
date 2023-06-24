package main

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
/*  暴力破解  时间复杂度 O(n3)
func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	l := len(s)
	longestStr := ""
	for i := 0; i < l; i++ {
		for j := i + 1; j < l+1; j++ {
			if isPalindrome(s[i:j]) {
				length := len(s[i:j])
				if length > len(longestStr) {
					longestStr = s[i:j]
				}
			}
		}
	}
	return longestStr
}
*/

// 中心扩散法
func longestPalindrome(s string) string {
	l := len(s)
	start, end := 0, 0
	for i := 0; i < l; i++ {
		//中间一个元素的扩散情况
		left1, right1 := expandAroundCenter(s, i, i)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		//中间两个元素的扩散情况
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func main() {
}
