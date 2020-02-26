package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		sub := target - num
		if v, ok := m[sub]; ok {
			return []int{v, i}
		} else {
			m[num] = i
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTwoSum(t *testing.T) {
	testFunc(t, twoSum)
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{11, 15, 2, 7}, 9)
	}
}

// 暴力法
func twoSum1(nums []int, target int) []int {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TestTwoSum1(t *testing.T) {
	testFunc(t, twoSum1)
}

func BenchmarkTwoSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum1([]int{11, 15, 2, 7}, 9)
	}
}

// 两遍哈希
func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		// 需要检测是不是同一个引用
		if v, ok := m[target-num]; ok && v != i {
			return []int{i, v}
		}
	}
	return nil
}

func TestTwoSum2(t *testing.T) {
	testFunc(t, twoSum2)
}

//testFunc has couple of test cases
func testFunc(t *testing.T, f func(nums []int, target int) []int) {
	ct := time.Now().Nanosecond()
	nums := []int{3, 2, 4}
	target := 6
	result := f(nums, target)
	assert.Equal(t, []int{1, 2}, result)
	nums = []int{2, 7, 11, 15}
	target = 9
	result = f(nums, target)
	assert.Equal(t, []int{0, 1}, result)
	t.Log(time.Now().Nanosecond() - ct)
}
