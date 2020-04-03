package cn

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。 
//
// 示例: 
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6 
// Related Topics 栈 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	l := list.New()
	sum := 0
	for i := 0; i < len(height); i++ {
		for l.Len() != 0 && height[l.Back().Value.(int)] < height[i] {
			h := height[l.Back().Value.(int)]
			l.Remove(l.Back())
			if l.Len() == 0 {
				break
			}
			distance := i - l.Back().Value.(int) - 1
			min := 0
			if height[l.Back().Value.(int)] < height[i] {
				min = height[l.Back().Value.(int)]
			} else {
				min = height[i]
			}
			sum += (min - h) * distance
		}
		l.PushBack(i)
	}
	return sum
}

func getMaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTrappingRainWater(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap_1(height []int) int {
	sum := 0
	// 从第一层开始到最高层遍历
	for i := 1; i <= getMax(height); i++ {
		// flag 用来决定是否有左边界
		flag := false
		// 小于当前高度的索引数
		tempSum := 0
		// 遍历数组所有的元素
		for j := 0; j < len(height); j++ {
			// 当前元素要小于外层高度，而且有左边界
			if height[j] < i && flag {
				tempSum++
			}
			// 只要有元素大于外层高度，就可以把tempSum加到结果中，同时确定有左边界
			if height[j] >= i {
				flag = true
				sum += tempSum
				tempSum = 0
			}
		}
	}
	return sum
}

func getMax(height []int) (max int) {
	for _, v := range height {
		if v > max {
			max = v
		}
	}
	return
}

// 按列求解
func trap_2(height []int) int {
	if len(height) < 1 {
		return 0
	}
	sum := 0
	// 第0和最后一列是不用求的，因为没有左/右边界
	for i := 1; i < len(height)-1; i++ {
		leftMax := 0
		rightMax := 0
		// 求左边最高的
		for j := i - 1; j >= 0; j-- {
			if leftMax < height[j] {
				leftMax = height[j]
			}
		}
		// 求右边最高的
		for j := i + 1; j < len(height); j++ {
			if rightMax < height[j] {
				rightMax = height[j]
			}
		}
		// 计算左右边界哪个小
		minB := 0
		if leftMax > rightMax {
			minB = rightMax
		} else {
			minB = leftMax
		}
		// 只有较小的一段大于当前列的高度才会有水，其他情况不会有水
		if minB > height[i] {
			sum += minB - height[i]
		}
	}
	return sum
}

// 动态规划
func trap_3(height []int) int {
	if len(height) < 1 {
		return 0
	}
	sum := 0
	leftMaxs := make([]int, len(height))
	rightMaxs := make([]int, len(height))
	// 右边最大列
	for i := 1; i < len(height)-1; i++ {
		leftMaxs[i] = getMaxInt(height[i-1], leftMaxs[i-1])
	}
	// 左边最大列
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxs[i] = getMaxInt(height[i+1], rightMaxs[i+1])
	}
	// 遍历每一列
	for i := 1; i < len(height)-1; i++ {
		minB := 0
		if rightMaxs[i] > leftMaxs[i] {
			minB = leftMaxs[i]
		} else {
			minB = rightMaxs[i]
		}
		if height[i] < minB {
			sum += minB - height[i]
		}
	}
	return sum
}

// 动态规划-优化左边
func trap_4(height []int) int {
	if len(height) < 1 {
		return 0
	}
	sum := 0
	leftMax := 0 // 左边最大的值
	rightMaxs := make([]int, len(height))
	// 左边最大列
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxs[i] = getMaxInt(height[i+1], rightMaxs[i+1])
	}
	// 遍历每一列
	for i := 1; i < len(height)-1; i++ {
		minB := 0
		leftMax = getMaxInt(height[i-1], leftMax)
		if rightMaxs[i] > leftMax {
			minB = leftMax
		} else {
			minB = rightMaxs[i]
		}
		if height[i] < minB {
			sum += minB - height[i]
		}
	}
	return sum
}

func trap_5(height []int) int {
	if len(height) < 2 {
		return 0
	}
	sum := 0
	leftMax := 0  // 左边最大的值
	rightMax := 0 // 右边最大的值
	leftIndex := 1
	rightIndex := len(height) - 2
	for leftIndex <= rightIndex {
		// 从左到右
		if height[leftIndex-1] < height[rightIndex+1] {
			leftMax = getMaxInt(leftMax, height[leftIndex-1])
			if leftMax > height[leftIndex] {
				sum += leftMax - height[leftIndex]
			}
			leftIndex++
		} else {
			// 从右到左
			rightMax = getMaxInt(rightMax, height[rightIndex+1])
			if rightMax > height[rightIndex] {
				sum += rightMax - height[rightIndex]
			}
			rightIndex--
		}
	}
	return sum
}
