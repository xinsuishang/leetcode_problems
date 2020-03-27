package hasGroupsSizeX

// 题目：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
// 给定一副牌，每张牌上都写着一个整数。
//此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：
//每组都有 X 张牌。
//组内所有的牌上都写着相同的整数。
//仅当你可选的 X >= 2 时返回 true。

// 解法
// 遍历计数，求最大公约数，最大公约数大于1的时候，true

func hasGroupsSizeX(deck []int) bool {
	dMap := make(map[int]int)
	for _, v := range deck {
		dMap[v]++
	}
	gCount := -1
	for _, count := range dMap {
		if gCount == -1 {
			gCount = count
		} else {
			gCount = gcd(gCount, count)
		}
	}
	return gCount > 1
}

func gcd(num1, num2 int) int {
	if num1 == 0 {
		return num2
	} else {
		return gcd(num2%num1, num1)
	}
}

// times: 1
