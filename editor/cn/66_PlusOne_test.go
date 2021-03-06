package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
// 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。 
//
// 你可以假设除了整数 0 之外，这个整数不会以零开头。 
//
// 示例 1: 
//
// 输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
// 
//
// 示例 2: 
//
// 输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。
// 
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)

// 数学计算法 - 优化
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		x := (digits[i] + 1) % 10
		digits[i] = x
		if x != 0 {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{8, 0, 0, 0}, plusOne([]int{7, 9, 9, 9}))
	assert.Equal(t, []int{1, 0, 0, 0, 0}, plusOne([]int{9, 9, 9, 9}))
	assert.Equal(t, []int{1, 2, 3}, plusOne([]int{1, 2, 2}))
}

// 数学计算法
func plusOne_1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		x := (digits[i] + 1) % 10
		if i == 0 && x == 0 {
			digits = make([]int, len(digits)+1)
			digits[0] = 1
			break
		}
		if x == 0 {
			digits[i] = 0
		} else {
			digits[i] = x
			break
		}
	}
	return digits
}
