package cn

//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例： 
//给定二叉树 [3,9,20,null,null,15,7]， 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回它的最大深度 3 。 
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root} // 存放每一层的元素
	res := 0
	for len(q) > 0 {
		curLen := len(q)
		// 遍历当前层的所有元素
		for i := 0; i < curLen; i++ {
			// 把下一层的所有元素放到队列中
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		// 更新队列，去掉当前层的元素
		q = q[curLen:]
		res++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	if l > r {
		return l
	}
	return r
}
