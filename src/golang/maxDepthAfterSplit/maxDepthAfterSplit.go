package maxDepthAfterSplit

// 1111. 有效括号的嵌套深度
// 题目：https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
// 给你一个「有效括号字符串」 seq，请你将其分成两个不相交的有效括号字符串，A 和 B，并使这两个字符串的深度最小。

// 解法
// 官方解释贼绕，看了题解理解就是把其中一部分（）分别替换成{}，求每个的嵌套层数最少的情况

func maxDepthAfterSplit(seq string) []int {
	length := len(seq)
	ans := make([]int, length, length)
	d := 0
	for idx, c := range seq {
		if c == '(' {
			ans[idx] = d % 2
			d += 1
		} else {
			d -= 1
			ans[idx] = d % 2
		}
	}
	return ans
}

// times: 1
