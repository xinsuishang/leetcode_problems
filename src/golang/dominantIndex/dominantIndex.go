package dominantIndex

// 题目：https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/
// 在一个给定的数组nums中，总是存在一个最大元素 。
// 查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
// 如果是，则返回最大元素的索引，否则返回-1

// 解法
// 遍历数组，找到最大的数和第二大的数，然后比较

func dominantIndex(nums []int) int {
	mIdx, max, second := 0, 0, 0
	for i, v := range nums {
		if v > second {
			second = v
		}
		if second > max {
			max, second = second, max
			mIdx = i
		}
	}
	if max >= second*2 {
		return mIdx
	} else {
		return -1
	}
}

// times: 1
