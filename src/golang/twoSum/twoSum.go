package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	// var numMap map[int]int
	var numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var difference = target - nums[i]

		if v, ok := numMap[difference]; ok {
			return []int{v, i}
		} else {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}

// times: 1
