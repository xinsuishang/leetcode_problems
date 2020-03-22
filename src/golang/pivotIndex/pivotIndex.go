package pivotIndex

// 题目：https://leetcode-cn.com/problems/find-pivot-index/
// 给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
// 我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
// 如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

// 解法
// 数组求和，从左向右遍历，满足sum-nums[i] - sumLeft = sumLeft时返回i，否则返回-1

func pivotIndex(nums []int) int {
	length := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	leftSum, i := 0, 0

	for {
		if i == length {
			return -1
		}
		if (sum - leftSum - nums[i]) == leftSum {
			return i
		}
		leftSum += nums[i]
		i++
	}
}

// times: 1
