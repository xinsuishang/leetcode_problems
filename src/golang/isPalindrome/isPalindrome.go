package isPalindrome

import "strings"

// 题目：https://leetcode-cn.com/problems/valid-palindrome/
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 解法
// 先转成小写，只保留有效信息到一个rune的slice内，然后双指针遍历
// 回文前后一致，条件为i<j

func isPalindrome(s string) bool {
	var check []rune
	for _, c := range strings.ToLower(s) {
		if 'a' <= c && c <= 'z' {
			check = append(check, c)
		} else if '0' <= c && c <= '9' {
			check = append(check, c)
		}
	}
	length := len(check)
	i, j := 0, length-1
	for length > 0 && j > -1 {
		if check[i] == check[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true

}

// times: 1
