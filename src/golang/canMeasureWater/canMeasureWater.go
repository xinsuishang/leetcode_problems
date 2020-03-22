package main

import (
	"container/list"
	"fmt"
)

// 题目：https://leetcode-cn.com/problems/water-and-jug-problem/
// 有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
//如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
//你允许：
//装满任意一个水壶
//清空任意一个水壶
//从一个水壶向另外一个水壶倒水，直到装满或者倒空

// 解法
// 最小0最大x+y，否则的话遍历所有可能，等于z返回
// v为上次剩余的答案；v+x<x+y, v+x为一个结果；v+y<x+y，v+y为一个结果;v-x > 0,v-x为一个结果;v-y > 0,v-y为一个结果

func canMeasureWater(x int, y int, z int) bool {

	if x == z || y == z || x+y == z {
		return true
	}
	if x+y < z || z < 0 {
		return false
	}
	rSet := make(map[int]int)
	l := list.List{}
	l.PushBack(0)
	for l.Len() != 0 {
		n := l.Front()
		l.Remove(n)
		v := n.Value.(int)
		if _, ok := rSet[v+x]; !ok && v+x < x+y {
			rSet[v+x] = 0
			l.PushBack(v + x)
		}
		if _, ok := rSet[v+y]; !ok && v+y < x+y {
			rSet[v+y] = 0
			l.PushBack(v + y)
		}
		if _, ok := rSet[v-x]; !ok && v-x >= 0 {
			rSet[v-x] = 0
			l.PushBack(v - x)
		}
		if _, ok := rSet[v-y]; !ok && v-y >= 0 {
			rSet[v-y] = 0
			l.PushBack(v - y)
		}
		if _, ok := rSet[z]; ok {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canMeasureWater(3, 4, 5))
}
