package isRectangleOverlap

// 题目：https://leetcode-cn.com/problems/rectangle-overlap/
// 矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。给出两个矩形，判断它们是否重叠并返回结果

// 解法
// 第二个矩形最大的值比第一个矩形最小的值小/最小的值比最大的值大就不满足。4个或之后取反。
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if len(rec1) == len(rec2) && len(rec1) == 4 {
		x1, y1, x2, y2 := rec1[0], rec1[1], rec1[2], rec1[3]
		r2x1, r2y1, r2x2, r2y2 := rec2[0], rec2[1], rec2[2], rec2[3]
		return !((x2 <= r2x1) || (x1 >= r2x2) || (y2 <= r2y1) || (y1 >= r2y2))
	}
	return false
}
