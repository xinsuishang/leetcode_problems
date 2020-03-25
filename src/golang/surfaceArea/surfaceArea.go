package surfaceArea

import "math"

// 题目：https://leetcode-cn.com/problems/surface-area-of-3d-shapes/
// 在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。
//每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。
//请你返回最终形体的表面积

// 解法
// 每个位置面积为4*height+2，多个正方体相接，会有两个面消除

func surfaceArea(grid [][]int) int {
	length, area := len(grid), 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			height := grid[i][j]
			if height > 0 {
				area += height*4 + 2
			}
			if i > 0 {
				area -= int(math.Min(float64(grid[i-1][j]), float64(grid[i][j]))) * 2
			}
			if j > 0 {
				area -= int(math.Min(float64(grid[i][j-1]), float64(grid[i][j]))) * 2
			}
		}
	}
	return area

}

// times: 1
