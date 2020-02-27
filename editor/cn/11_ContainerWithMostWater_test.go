package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。 
//
// 
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。 
//
// 
//
// 示例: 
//
// 输入: [1,8,6,2,5,4,8,3,7]
//输出: 49 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	result := 0
	lpt := 0
	rpt := len(height) - 1
	for lpt < rpt {
		cur := 0
		if height[lpt] < height[rpt] {
			cur = height[lpt] * (rpt - lpt)
			lpt++
		} else {
			cur = height[rpt] * (rpt - lpt)
			rpt--
		}
		if cur > result {
			result = cur
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {
	testContainerWithMostWater(t, maxArea1)
	testContainerWithMostWater(t, maxArea)
}

func BenchmarkTestContainerWithMostWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{1, 3, 2, 5, 25, 24, 5}
		//maxArea1(nums) // 26.2 ns/op
		maxArea(nums) // 8.34 ns/op
	}
}

func testContainerWithMostWater(t *testing.T, maxArea func(height []int) int) {
	nums := []int{1, 3, 2, 5, 25, 24, 5}
	target := 24
	assert.Equal(t, maxArea(nums), target)
	nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	target = 49
	assert.Equal(t, maxArea(nums), target)
}

// 暴力求解
func maxArea1(height []int) int {
	result := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			cur := 0
			if height[i] > height[j] {
				cur = height[j] * (j - i)
			} else {
				cur = height[i] * (j - i)
			}
			if cur > result {
				result = cur
			}
		}
	}
	return result
}
