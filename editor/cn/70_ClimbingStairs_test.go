package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ 
//
// 注意：给定 n 是一个正整数。 
//
// 示例 1： 
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶 
//
// 示例 2： 
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
// 
// Related Topics 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
// 利用切片
func climbStairs(n int) int {
	list := make([]int, n+2)
	list[0] = 1 // f(0)
	list[1] = 1 // f(1)
	for i := 2; i <= n; i++ {
		list[i] = list[i-1] + list[i-2] // f(n)
	}
	return list[n]
}

//leetcode submit region end(Prohibit modification and deletion)

/*
分析：最后只能走1或者2步
f(1) = 1
f(2) = 2
f(3) = f(2) + f(1)
f(4) = f(3) + f(2)
f(5) = f(4) + f(3)
====> f(n) = f(n-1) + f(n-2)
*/

func TestClimbingStairs(t *testing.T) {
	assert.Equal(t, climbStairs(5), 8)
	assert.Equal(t, climbForce(5, 0), 8)
	assert.Equal(t, climbStairs1(5), 8)
}

func BenchmarkClimbingStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//climbForce(10,0) // 488 ns/op
		//climbStairs1(10) // 3.93 ns/op
		climbStairs(10) // 38.0 ns/op
	}
}

// 暴力法
func climbForce(n int, cur int) int {
	if n < cur {
		return 0
	}

	if n == cur {
		return 1
	}
	return climbForce(n, cur+1) + climbForce(n, cur+2)
}

// 斐波那公式 f(n) = f(n-1) + f(n-2)
func climbStairs1(n int) int {
	//return climbForce(n, 0)
	if n == 1 || n == 0{
		return 1
	}
	if n == 2 {
		return 2
	}
	s1 := 1
	s2 := 2
	re := 0
	for i := 3; i <= n; i++ {
		re = s1 + s2
		s1, s2 = s2, re
	}
	return re
}
