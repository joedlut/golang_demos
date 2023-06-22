package main

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//滑动窗口

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		//向右移动做指针
		if i != 0 {
			delete(m, s[i-1])
		}
		//如果字符集合中没有重复就继续移动右指针
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
}
