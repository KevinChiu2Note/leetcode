package cn

import "container/list"

//给定一个 N 叉树，返回其节点值的前序遍历。
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 返回其前序遍历: [1,3,5,6,2,4]。 
//
// 
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 使用栈
func preorder(root *Node) []int {
	var ret []int
	if root == nil {
		return ret
	}
	l := list.New()
	l.PushFront(root)
	for l.Len() > 0 {
		f := l.Front().Value.(*Node)
		ret = append(ret, f.Val)
		l.Remove(l.Front())
		for i := len(f.Children) - 1; i >= 0; i-- {
			l.PushFront(f.Children[i])
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
// 递归法
func preorder_1(root *Node) []int {
	var ret []int
	if root != nil {
		ret = append(ret, root.Val)
		for _, n := range root.Children {
			ret = append(ret, preorder(n)...)
		}
	}
	return ret
}
