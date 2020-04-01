package cn

import (
	"container/list"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//
// 返回滑动窗口中的最大值。 
//
// 
//
// 进阶： 
//
// 你能在线性时间复杂度内解决此题吗？ 
//
// 
//
// 示例: 
//
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7] 
//解释: 
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10^5 
// -10^4 <= nums[i] <= 10^4 
// 1 <= k <= nums.length 
// 
// Related Topics 堆 Sliding Window

//leetcode submit region begin(Prohibit modification and deletion)
// 双端队列
func maxSlidingWindow(nums []int, k int) []int {
	// 思路：
	//1. 遍历元素 nums[i] ，然后跟队列尾部元素比较，如果比尾部元素大，就出队，然后继续比较，直到 nums[i] 小于尾部元素，然后将它入队。
	//2. 然后用一下队列首部元素的下标，计算出队列中区间的长度，如果大于 k 了，那么队首元素就要出队。
	//3. 最后队首元素就是当前区间的最大值。
	if len(nums) == 0 {
		return nums
	}
	l := list.New()
	result := make([]int, len(nums)-k+1)
	index := 0
	for i := 0; i < len(nums); i++ {
		for l.Len() != 0 && nums[i] >= nums[l.Back().Value.(int)] {
			l.Remove(l.Back())
		}
		l.PushBack(i)
		if i-l.Front().Value.(int)+1 > k {
			l.Remove(l.Front())
		}
		if i >= k-1 {
			result[index] = nums[l.Front().Value.(int)]
			index++
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSlidingWindowMaximum(t *testing.T) {
	fmt.Printf("%T\n", zero)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

// 暴力法
func maxSlidingWindow_1(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	result := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums)+1-k; i++ {
		curMax := math.MinInt32
		for j := i; j < i+k; j++ {
			if nums[j] > curMax {
				curMax = nums[j]
			}
		}
		result = append(result, curMax)
	}
	return result
}

const zero = 0.0

// 指针法
func maxSlidingWindow_2(nums []int, k int) []int {
	// 判断数组长度，不符合直接返回
	if len(nums) == 0 {
		return nums
	}
	// 整个数组最大值的索引
	maxPos := -1
	// 返回的数组
	result := make([]int, len(nums)-k+1)
	// i 只需要从0到len(nums)-k，len(nums)-k 后面的元素将由下面的for进行遍历
	for i := 0; i <= len(nums)-k; i++ {
		// j 为窗口的右边界
		j := i + k - 1
		// 最大值的索引不是-1 并且 右边界大于全局最大值
		if maxPos != -1 && nums[j] >= nums[maxPos] {
			maxPos = j
			result[i] = nums[maxPos]
		} else if i <= maxPos { // 如果左边界小于最大值的索引那么最大值不变
			result[i] = nums[maxPos]
		} else { // 最大值的索引不在窗口中，需要重新确定最大值和其索引
			curMax := math.MinInt32
			curMaxPos := 0
			for z := i; z <= j; z++ {
				if nums[z] > curMax {
					curMaxPos = z
					curMax = nums[z]
				}
			}
			maxPos = curMaxPos
			result[i] = nums[maxPos]
		}
	}
	return result
}

// 双端队列 - 不好理解
func maxSlidingWindow_3(nums []int, k int) []int {
	// 判断数组长度，不符合直接返回
	if len(nums) == 0 {
		return nums
	}
	l := list.New()
	result := make([]int, len(nums)-k+1)
	index := 0
	for i := 0; i < len(nums); i++ {
		// 如果i>=k，说明窗口已经初始化完成了
		if i >= k {
			// 队列中的第一个值就是
			if l.Front().Value.(int) == nums[i-k] {
				l.Remove(l.Front())
			}
		}
		for l.Len() != 0 && nums[i] > l.Back().Value.(int) {
			l.Remove(l.Back())
		}
		l.PushBack(nums[i])
		if i >= k-1 {
			result[index] = l.Front().Value.(int)
			index++
		}
	}
	return result
}
