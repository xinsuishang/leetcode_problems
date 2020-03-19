package longestPalindrome

// 题目：https://leetcode-cn.com/problems/longest-palindrome/
// 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

// 解法
// 统计所有给出的字母出现的次数，回文串字母除最中间的一个外都出现多次，取每个出现的最大偶数次，如果有奇数，加一

func longestPalindrome(s string) int {
	count := 0
	dict := make(map[rune]int)
	for _, c := range s {
		dict[c]++
	}
	for _, v := range dict {
		count += v / 2 * 2
	}
	if count != len(s) {
		count++
	}
	return count
}

// times: 1
