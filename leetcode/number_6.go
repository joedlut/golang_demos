package main

/*
6. N 字形变换
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);


示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I
示例 3：

输入：s = "A", numRows = 1
输出："A"


提示：

1 <= s.length <= 1000
s 由英文字母（小写和大写）、',' 和 '.' 组成
1 <= numRows <= 1000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

这个题包装的很唬人，但是实际就是给s的每一位标记就好了：假设numRows为4，那就是那s每一位的行数就是：1234321234321
*/

func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if n == 1 || r >= n || r == 1 {
		return s
	}
	//周期数
	t := 2*r - 2
	//列数  n/t*(r-1) 向上取整
	c := (n + t - 1) / t * (r - 1)
	matrix := make([][]byte, r)
	for i := 0; i < r; i++ {
		matrix[i] = make([]byte, c)
	}
	x, y := 0, 0
	for j := 0; j < n; j++ {
		matrix[x][y] = byte(s[j])
		//
		if j%t < r-1 {
			x++ //向下移动
		} else {
			//向右上移动
			x--
			y++
		}
	}
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] != 0 {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return string(ans)

}
