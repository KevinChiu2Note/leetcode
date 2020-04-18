package cn

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 
//
// 示例： 
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
// 
// Related Topics 字符串 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	// 使用new方法开辟内存返回内存地址
	res := new([]string)
	dfs(n, n, "", res)
	return *res
}

func dfs(l int, r int, str string, res *[]string) {
	/*
	   回溯跳出条件，
	   并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	   所以右括号用完了就意味着左括号必定用完了
	*/
	if r == 0 {
		*res = append(*res, str)
		return
	}
	// 生成左括号
	if l > 0 {
		dfs(l-1, r, str+"(", res)
	}
	// 括号成对存在，有左括号才会有右括号
	if r > l {
		dfs(l, r-1, str+")", res)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
func TestGenerateParenthesis(t *testing.T) {
	Convey("测试括号生成", t, func() {
		res := generateParenthesis(1)
		So(res, ShouldResemble, []string{"(())", "()()"})
	})
}
