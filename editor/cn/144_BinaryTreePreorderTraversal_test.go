package cn

//给定一个二叉树，返回它的 前序 遍历。
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [1,2,3]
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代法-Morris 保持结构
func preorderTraversal(root *TreeNode) []int {
	var max *TreeNode
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}
			if max.Right == nil {
				res = append(res, root.Val)
				max.Right = root
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
// 递归法
func preorderTraversal_1(root *TreeNode) []int {
	var ret []int
	if root != nil {
		ret = append(ret, root.Val)
		ret = append(ret, preorderTraversal(root.Left)...)
		ret = append(ret, preorderTraversal(root.Right)...)
	}
	return ret
}

// 迭代法-使用栈
func preorderTraversal_2(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		pre := len(stack) - 1
		root = stack[pre].Right
		stack = stack[:pre]
	}
	return ret
}
