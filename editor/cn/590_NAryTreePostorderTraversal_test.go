package cn

//给定一个 N 叉树，返回其节点值的后序遍历。
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 返回其后序遍历: [5,6,3,2,4,1]. 
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

type Node struct {
	Val      int
	Children []*Node
}

// 迭代法 - 使用栈
func postorder(root *Node) []int {
	var tmp []int
	if root == nil {
		return tmp
	}
	stack := make([]*Node, 0)
	// root元素入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 只要栈不为空就弹栈，并把值放到缓存结果中
		n := stack[len(stack)-1]
		tmp = append(tmp, n.Val)
		stack = stack[:len(stack)-1]
		// 遍历root元素的子元素，全部入栈
		if len(n.Children) > 0 {
			stack = append(stack, n.Children...)
		}
	}
	// 得到缓存结果后，把缓存结果反转
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}

//leetcode submit region end(Prohibit modification and deletion)
// 递归法
func postorder_1(root *Node) []int {
	var ret []int
	if root != nil {
		for _, child := range root.Children {
			ret = append(ret, postorder(child)...)
		}
		ret = append(ret, root.Val)
	}
	return ret
}
