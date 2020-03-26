package cn

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。
//
// 
//
// 说明: 
//
// 
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。 
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。 
// 
//
// 
//
// 示例: 
//
// 输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6] 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
// 多个指针 - 后往前 - 优化
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeSortedArray(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
	nums3 := []int{4, 5, 6, 0, 0, 0}
	nums4 := []int{1, 2, 3}
	merge(nums3, 3, nums4, 3)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums3)

}

// 使用排序
func merge_1(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

// 使用额外数组和多个指针 - 前往后
func merge_2(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, len(nums1))
	copy(nums1Copy, nums1)
	var p1, p2, p int
	for m > p1 && n > p2 {
		if nums1Copy[p1] > nums2[p2] {
			nums1[p] = nums2[p2]
			p2++
		} else {
			nums1[p] = nums1Copy[p1]
			p1++
		}
		p++
	}
	if m > p1 {
		nums1 = append(nums1[:p], nums1Copy[p1:m]...)
	}
	if n > p2 {
		nums1 = append(nums1[:p], nums2[p2:]...)
	}
}

// 多个指针 - 后往前
func merge_3(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	copy(nums1, nums2[:p2+1]) // 最后p2的边界需要注意
}
