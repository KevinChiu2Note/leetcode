package cn

import (
	"container/list"
	"fmt"
	"testing"
)

//翻转一棵二叉树。
//
// 示例： 
//
// 输入： 
//
//      4
//   /   \
//  2     7
// / \   / \
//1   3 6   9 
//
// 输出： 
//
//      4
//   /   \
//  7     2
// / \   / \
//9   6 3   1 
//
// 备注: 
//这个问题是受到 Max Howell 的 原问题 启发的 ： 
//
// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。 
// Related Topics 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代法
func invertTree(root *TreeNode) *TreeNode {
	l := list.New()
	// 特判
	if root != nil {
		l.PushBack(root)
	}
	for l.Len() != 0 {
		// 出栈
		node := l.Back().Value.(*TreeNode)
		l.Remove(l.Back())
		// 压栈
		if node.Right != nil {
			l.PushBack(node.Right)
		}
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		// 交换
		node.Right, node.Left = node.Left, node.Right
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
// 递归法
func invertTree_1(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree_1(root.Left)
		invertTree_1(root.Right)
	}
	return root
}

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 7}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 6}
	node6 := &TreeNode{Val: 9}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6

	tree := invertTree(root)
	fmt.Println(tr(tree))

}

func tr(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, root.Val)
		res = append(res, tr(root.Right)...)
		res = append(res, tr(root.Left)...)
	}
	return res
}
