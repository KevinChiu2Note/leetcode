package cn

//给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 返回其层序遍历: 
//
// [
//     [1],
//     [3,2,4],
//     [5,6]
//]
// 
//
// 
//
// 说明: 
//
// 
// 树的深度不会超过 1000。 
// 树的节点总数不会超过 5000。 
// Related Topics 树 广度优先搜索

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 递归法
func levelOrder(root *Node) [][]int {
	var res [][]int
	recursive(root, 0, &res)
	return res
}

func recursive(root *Node, level int, res *[][]int) {
	if root == nil {
		return
	}
	// 遍历子节点
	for _, v := range root.Children {
		recursive(v, level+1, res)
	}
	// 当前层要和结果层数要相等，否则添加结果时空指针异常
	for len(*res) <= level {
		*res = append(*res, []int{})
	}
	// 结果添加
	(*res)[level] = append((*res)[level], root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
func levelOrder_1(root *Node) [][]int {
	var res [][]int
	var arr []*Node
	if root != nil {
		arr = append(arr, root)
		for len(arr) > 0 {
			temp := len(arr)
			var level []int
			for i := 0; i < temp; i++ {
				level = append(level, arr[i].Val)
				if len(arr[i].Children) > 0 {
					arr = append(arr, arr[i].Children...)
				}
			}
			res = append(res, level)
			arr = arr[temp:]
		}
	}
	return res
}

func levelOrder_2(root *Node) [][]int {
	var res [][]int
	// 数组用于保存每一层的元素
	var arr []*Node
	if root != nil {
		// root 元素放进数组中
		arr = append(arr, root)
		// 只要数组不为空就需要遍历
		for len(arr) != 0 {
			// 记录当前层的节点树
			temp := len(arr)
			// 每一层的节点的值
			var level []int
			// 遍历当前层的元素
			for i := 0; i < temp; i++ {
				level = append(level, arr[i].Val)
				// 把所有当前层的子节点放入数组中
				for _, v := range arr[i].Children {
					arr = append(arr, v)
				}
			}
			// 当前层遍历完了保存数据
			res = append(res, level)
			// 更新下一层要遍历的数组
			arr = arr[temp:]
		}
	}
	return res
}
