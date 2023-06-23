package main

import "fmt"

func generateAll(str string, n int) (strList []string) {
	if len(str) == n {
		if Validate(str) {
			strList = append(strList, str)
			fmt.Println(str)
		}
		return strList
	}
	str += "("
	generateAll(str, n)
	str = str[:len(str)-1]
	str += ")"
	generateAll(str, n)
	return nil
}

func Match(m, n byte) bool {
	if m == '(' && n == ')' {
		return true
	}
	return false
}

func Validate(s string) bool {
	length := len(s)
	if length == 1 {
		return false
	}
	if s[0] == ')' {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			top := len(stack) - 1
			if top < 0 {
				return false
			}
			if Match(stack[top], s[i]) {
				//pop stack
				stack = stack[:top]
				// 进入下一次循环
				continue
			}
		}
		stack = append(stack, s[i])
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	str := ""
	fmt.Println(generateAll(str, 6))
}
