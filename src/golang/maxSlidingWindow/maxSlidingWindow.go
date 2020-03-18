package maxSlidingWindow

// 题目：https://leetcode-cn.com/problems/sliding-window-maximum/
// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//返回滑动窗口中的最大值。

// 解法
// 切片作为双端队列，遍历数组，队列中存数组下标，队首为窗口可视范围内的最大值所在数组下标
// 遍历到的idx为窗口最右端元素

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var deque []int
	for idx, num := range nums {
		//如果队列中元素超过窗口可视范围，leftpop，保证队首在窗口可视范围内
		for len(deque) > 0 && deque[0] <= idx-k {
			deque = deque[1:]
		}
		//如果队尾元素小于当前元素，pop，保证队首为可视范围内最大值
		for len(deque) > 0 && num > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		//当前元素放入队列
		deque = append(deque, idx)
		//满足窗口条件，开始记录窗口内最大值
		if idx >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res

}

// times: 1
