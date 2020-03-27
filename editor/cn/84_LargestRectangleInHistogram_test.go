package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 
//
// 
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。 
//
// 
//
// 
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。 
//
// 
//
// 示例: 
//
// 输入: [2,1,5,6,2,3]
//输出: 10 
// Related Topics 栈 数组

//leetcode submit region begin(Prohibit modification and deletion)
// 栈
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 1, len(heights)+1)
	// 加入哨兵，不需要遍历完数组再循环处理栈
	heights = append(heights, -1)
	stack[0] = -1
	max := 0
	for i, height := range heights {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] > height {
			h := heights[stack[len(stack)-1]] // 要先拿到栈顶的高度，用于计算面积
			stack = stack[:len(stack)-1]      // 弹栈后，用于获得矩形的最大宽度
			max = maxNum(max, h*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
//leetcode submit region end(Prohibit modification and deletion)

func TestLargestRectangleInHistogram(t *testing.T) {
	assert.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

// 暴力法 - 超时
func largestRectangleArea_1(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			h := int(^uint(0) >> 1)
			for k := i; k <= j; k++ {
				if heights[k] < h {
					h = heights[k]
				}
			}
			tmp := h * (j - i + 1)
			if tmp > maxArea {
				maxArea = tmp
			}
		}
	}
	return maxArea
}

// 暴力法 - 优化
func largestRectangleArea_2(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		h := int(^uint(0) >> 1)
		for j := i; j < len(heights); j++ {
			if heights[j] < h {
				h = heights[j]
			}
			tmp := h * (j - i + 1)
			if tmp > maxArea {
				maxArea = tmp
			}
		}
	}
	return maxArea
}

// 分治
func largestRectangleArea_3(heights []int) int {
	return calculateArea(heights, 0, len(heights)-1)
}

func calculateArea(heights []int, start, end int) int {
	if start > end {
		return 0
	}
	// 寻找最小的值的索引
	shortIndex := start
	for i := start; i <= end; i++ {
		if heights[i] < heights[shortIndex] {
			shortIndex = i
		}
	}
	// 计算左边的区域
	left := calculateArea(heights, start, shortIndex-1)
	// 计算右边的区域
	right := calculateArea(heights, shortIndex+1, end)
	// 计算当前区域的面积
	cur := heights[shortIndex] * (end - start + 1)
	return maxNum(cur, maxNum(left, right))
}

