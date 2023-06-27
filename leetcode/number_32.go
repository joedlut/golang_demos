package main

import "fmt"

/*
32. 最长有效括号
困难
2.3K
company
亚马逊
company
Facebook
company
美团
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0
提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'
*/
func max1(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

/*
对于遇到的每个 ‘(’\text{‘(’}‘(’ ，我们将它的下标放入栈中
对于遇到的每个 ‘)’\text{‘)’}‘)’ ，我们先弹出栈顶元素表示匹配了当前右括号：
如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」

需要注意的是，如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，我们在一开始的时候往栈中放入一个值为 −1-1−1 的元素。

作者：力扣官方题解
链接：https://leetcode.cn/problems/longest-valid-parentheses/solutions/314683/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

作者：力扣官方题解
链接：https://leetcode.cn/problems/longest-valid-parentheses/solutions/314683/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
go for it !
*/
func longestValidParentheses(s string) int {
	stack := []int{}
	stack = append(stack, -1)
	maxLen := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			//pop the stack
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				//栈顶存的最后一个不匹配的右括号的下标, 或者下标-1
				// ()(())  5-(-1) = 6
				// )()()  4- 0 = 4
				// )())() 5 -3 = 2
				stack = append(stack, i)
			} else {
				maxLen = max1(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestValidParentheses("()((()))"))
}
