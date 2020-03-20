package main

import (
	"container/heap"
	"fmt"
)

// 题目：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

// 解法
// 1.排序，取切片内的数
// 2.topK，大顶堆

type intHeap []int

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	maxIdx := h.Len() - 1
	popValue := (*h)[maxIdx]
	*h = (*h)[:maxIdx]
	return popValue
}

func (h *intHeap) Peek() int {
	if h.Len() == 0 {
		return 0
	}
	return (*h)[0]
}

func getLeastNumbers(arr []int, k int) []int {
	var res []int
	if k == 0 {
		return res
	}
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	for _, v := range arr {
		if maxHeap.Len() < k {
			heap.Push(maxHeap, v)
		} else if maxHeap.Peek() > v {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, v)
		}
	}
	for maxHeap.Len() != 0 {
		res = append(res, heap.Pop(maxHeap).(int))
	}

	return res
}

func main() {
	arr := []int{3, 2, 1}
	k := 2
	fmt.Println(getLeastNumbers(arr, k))
}

// times: 1
