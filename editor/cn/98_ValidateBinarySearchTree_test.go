package cn

import (
	"math"
)

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征： 
//
// 
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
// 
// Related Topics 树 深度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 队列
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	qNode := []*TreeNode{root}
	qMax := []int{1<<63 - 1}
	qMin := []int{-1 << 63}
	for len(qNode) > 0 {
		max, min, node := qMax[0], qMin[0], qNode[0]
		if node.Val <= min || node.Val >= max {
			return false
		}
		if node.Left != nil {
			qNode = append(qNode, node.Left)
			qMin = append(qMin, min)
			qMax = append(qMax, node.Val)
		}
		if node.Right != nil {
			qNode = append(qNode, node.Right)
			qMin = append(qMin, node.Val)
			qMax = append(qMax, max)
		}
		qNode, qMax, qMin = qNode[1:], qMax[1:], qMin[1:]
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
// 前序递归
func isValidBST_1(root *TreeNode) bool {
	return DfsisValidBST(root, math.MaxInt64, math.MinInt64)
}

func DfsisValidBST(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	if !DfsisValidBST(root.Left, root.Val, min) {
		return false
	}
	if !DfsisValidBST(root.Right, max, root.Val) {
		return false
	}
	return true
}

// 前序迭代
func isValidBST_2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root}
	minS := []int{-1 << 63}
	maxS := []int{1<<63 - 1}
	for len(stack) > 0 {
		pre := len(stack) - 1
		cur, min, max := stack[pre], minS[pre], maxS[pre]
		stack, minS, maxS = stack[:pre], minS[:pre], maxS[:pre]
		for cur != nil {
			if cur.Val <= min || cur.Val >= max {
				return false
			}
			// 右边子节点压入栈中
			maxS = append(maxS, max)
			// 右边子节点的最小值
			minS = append(minS, cur.Val)
			// 右边子节点的最大值
			stack = append(stack, cur.Right)
			// 设置左边节点不能大于max
			max = cur.Val
			// 下一个左节点
			cur = cur.Left
		}
	}
	return true
}

// 中序递归
var last *TreeNode

func isValidBST_3(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return dfs2(root)
}
func dfs2(node *TreeNode) bool {
	if node == nil {
		return true
	}
	// 检查左边的元素先 || 右边元素检查
	if !dfs2(node.Left) || node.Val <= last.Val {
		return false
	}
	// 右边元素要大于上一个元素
	last = node
	// 检查右边
	return dfs2(node.Right)
}

// 中序迭代
func isValidBST_4(root *TreeNode) bool {
	var stack []*TreeNode
	temp := &TreeNode{Val: -1 << 63}
	for len(stack) > 0 || root != nil {
		// 压入所有左边的
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出
		pre := len(stack) - 1
		// 检测
		if stack[pre].Val <= temp.Val {
			return false
		}
		// 确认新的最小值
		temp = stack[pre]
		root = stack[pre].Right
		stack = stack[:pre]
	}
	return true
}
