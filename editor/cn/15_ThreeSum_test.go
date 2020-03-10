package cn

import (
	"sort"
	"strconv"
	"testing"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三
//元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例： 
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
// 夹逼法-优化 40ms 6.9MB
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// 排序
	sort.Ints(nums)
	// 结果列表保存
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, left, right := -nums[i], i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if target == sum {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// 去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 去重
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if target < sum {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeSum(t *testing.T) {
	//nums := []int{4, -1, 0, 1, 2, -1}
	//nums := []int{0, 0, 0, 0}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-2,0,1,1,2}
	sum := threeSum(nums)
	t.Log(sum)
}

// 暴力法 - 不能去重，时间复杂度：O(n^3)，空间复杂度：O(1)
func threeSum_Force(nums []int) [][]int {
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] == nums[j]+nums[k] {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

// 暴力法 - 去重，时间复杂度：O(n^3)，空间复杂度：O(n)使用了map 超时
func threeSum_Force2(nums []int) [][]int {
	// 一定要先排序
	sort.Ints(nums) // 快排
	var result [][]int
	m := map[string]int{}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					str := strconv.Itoa(nums[i]) + "_" + strconv.Itoa(nums[j]) + "_" + strconv.Itoa(nums[k])
					if _, ok := m[str]; ok {
						continue
					}
					result = append(result, []int{nums[i], nums[j], nums[k]})
					m[str] = 0
				}
			}
		}
	}
	return result
}

// 哈希表-慢 500ms 35.6MB
func threeSum_Hash(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// 排序
	sort.Ints(nums)
	// 结果列表保存
	var res [][]int
	// 结果去重
	resMap := map[string]int{}
	// 循环到倒数第3个即可，因为最后3个数字可能成为唯一的解
	for i := 0; i < len(nums)-2; i++ {
		// 重复的目标略过
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		// map 给 twoSum 使用
		m := make(map[int]int)
		target := -nums[i]
		// twoSum 的解法
		for j := i + 1; j < len(nums); j++ {
			diff := target - nums[j]
			if _, ok := m[diff]; ok {
				tmp := []int{nums[i], nums[j], diff}
				sort.Ints(tmp)
				str := strconv.Itoa(tmp[0]) + "_" + strconv.Itoa(tmp[1]) + "_" + strconv.Itoa(tmp[2])
				_, ok := resMap[str]
				if ok {
					continue
				}
				res = append(res, tmp)
				resMap[str] = 0
			} else {
				m[nums[j]] = nums[j]
			}
		}
	}
	return res
}

// 夹逼法-慢 432ms 8.5MB
func threeSum_twoSize(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// 排序
	sort.Ints(nums)
	// 结果列表保存
	var res [][]int
	// 结果去重
	resMap := map[string]int{}
	for i := 0; i < len(nums)-2; i++ {
		head := i + 1
		tail := len(nums) - 1
		for head < tail {
			sum := nums[head] + nums[tail]
			if sum == -nums[i] {
				str := strconv.Itoa(nums[i]) + "_" + strconv.Itoa(nums[head]) + "_" + strconv.Itoa(nums[tail])
				if _, ok := resMap[str]; ok {
					goto JUMP
				}
				res = append(res, []int{nums[i], nums[head], nums[tail]})
				resMap[str] = 0
			}
		JUMP:
			if -nums[i] <= sum {
				tail--
			} else {
				head++
			}
		}
	}
	return res
}
