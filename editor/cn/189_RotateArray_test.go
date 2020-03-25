package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
// 示例 1: 
//
// 输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
// 
//
// 示例 2: 
//
// 输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释: 
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100] 
//
// 说明: 
//
// 
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。 
// 要求使用空间复杂度为 O(1) 的 原地 算法。 
// 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
// 使用反转
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverseArray(nums, 0, len(nums)-1)
	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)
}

func reverseArray(list []int, start, end int) {
	for start < end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRotateArray(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(input, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, input)
}

// 暴力法
func rotate_1(nums []int, k int) {
	var tmp, pre int
	for i := k; i > 0; i-- {
		pre = nums[len(nums)-1] // 拿到修改前当前数组的最后那个数字
		for j := 0; j < len(nums); j++ {
			// 开始交换
			tmp = nums[j]
			nums[j] = pre
			pre = tmp
		}
	}
}

// 使用数组
func rotate_2(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	for i, v := range tmp {
		nums[i] = v
	}
}

// 使用数组优化
func rotate_21(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, tmp)
}

// 环状交替
func rotate_3(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		curIndex := start
		preValue := nums[start]
		for {
			count++
			nextIndex := (k + curIndex) % len(nums)
			//nextValue := nums[nextIndex]
			// 交换
			//tmp:= nextValue
			//nums[nextIndex] = preValue
			//preValue = tmp
			nums[nextIndex], preValue = preValue, nums[nextIndex]
			curIndex = nextIndex
			if start == curIndex {
				break
			}
		}
	}
}
