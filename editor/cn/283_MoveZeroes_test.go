package cn

import (
	"testing"
)

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例: 
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0] 
//
// 说明: 
//
// 
// 必须在原数组上操作，不能拷贝额外的数组。 
// 尽量减少操作次数。 
// 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cur], nums[i] = nums[i], nums[cur]
			cur++
		}
	}
	//fmt.Println(nums)
}

//leetcode submit region end(Prohibit modification and deletion)

func BenchmarkMoveZeroes(b *testing.B) {
	nums := []int{0, 1, 0, 3, 12}
	for i := 0; i < b.N; i++ {
		//moveZeroes1(nums) // 8.69 ns/op
		//moveZeroes2(nums) // 7。25 ns/op
		moveZeroes(nums) // 5。89 ns/op
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

// 双指针，填充0
func moveZeroes1(nums []int) {
	count := 0
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// 双指针，统计0的个数
			count++
			continue
		}
		nums[cur] = nums[i]
		cur++
	}
	// 尾部填充0
	for i := 0; i < count; i++ {
		nums[cur] = 0
		cur++
	}
}

// 双指针，填充0，优化版
func moveZeroes2(nums []int) {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nums[cur] = nums[i]
		cur++
	}
	for ; cur < len(nums); cur++ {
		nums[cur] = 0
	}
}
