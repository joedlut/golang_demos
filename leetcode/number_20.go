package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isMatch(left byte, right byte) bool {
	if left == '{' && right == '}' {
		return true
	}
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	return false
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	/*
		if s[0] == '}' || s[0] == ']' || s[0] == ')' {
			return false
		}*/
	stack := make([]byte, 0)
	for i := 0; i <= len(s)-1; i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			l := len(stack)
			//要考虑下标为0的场景
			if l > 0 {
				top := stack[l-1]
				if isMatch(top, s[i]) {
					stack = stack[:l-1]
					continue
				}
			}
			return false
		}
	}
	//循环结束后记得判断栈的长度
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
