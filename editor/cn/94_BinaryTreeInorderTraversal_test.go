package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定一个二叉树，返回它的中序 遍历。
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
//输出: [1,3,2] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 哈希表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 莫尼斯 破坏结构
func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var max *TreeNode
	for root != nil {
		if root.Left == nil { // 左边没有子树，把当前节点数据输出
			ret = append(ret, root.Val)
			root = root.Right
		} else { // 左边有子树
			max = root.Left // 索引，用来找到最深处的右子节点
			for max.Right != nil {
				max = max.Right
			}
			max.Right = root // 找到的节点的右边和root节点连接
			//root, root.Left = root.Left, nil // 下面3行的简写
			lTmp := root.Left // lTmp 就是root节点的左子节点
			root.Left = nil   // 砍掉左边的子树
			root = lTmp       // root节点为lTmp
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestBinaryTreeInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 3}
	root.Right = t1
	t1.Left = t2
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(root))
	root = &TreeNode{Val: 1}
	t1 = &TreeNode{Val: 2}
	t2 = &TreeNode{Val: 3}

	root.Left = t1
	root.Right = t2
	assert.Equal(t, []int{2, 1, 3}, inorderTraversal(root))
}

// 递归法
func inorderTraversal_1(root *TreeNode) []int {
	var xs []int
	if root != nil {
		xs = append(xs, inorderTraversal(root.Left)...)
		xs = append(xs, root.Val)
		xs = append(xs, inorderTraversal(root.Right)...)
	}
	return xs
}

// 迭代法 1
func inorderTraversal_2(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

// 迭代法 2
func inorderTraversal_3(root *TreeNode) []int {
	var ret []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 入栈
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			continue
		}
		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}
