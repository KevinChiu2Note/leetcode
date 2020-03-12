package cn

import "testing"

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。 
//
// 说明：不允许修改给定的链表。 
//
// 
//
// 示例 1： 
//
// 输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。
// 
//
// 
//
// 示例 2： 
//
// 输入：head = [1,2], pos = 0
//输出：tail connects to node index 0
//解释：链表中有一个环，其尾部连接到第一个节点。
// 
//
// 
//
// 示例 3： 
//
// 输入：head = [1], pos = -1
//输出：no cycle
//解释：链表中没有环。
// 
//
// 
//
// 
//
// 进阶： 
//你是否可以不用额外空间解决此题？ 
// Related Topics 链表 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// floyd 算法，快慢指针 - 优化
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLinkedListCycleIi(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node2
	result := detectCycle(node1)
	t.Log(result)
}

func detectCycle_map(head *ListNode) *ListNode {
	m := map[*ListNode]int{}
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = 1
		head = head.Next
	}
	return nil
}

// floyd 算法，快慢指针
func detectCycle_1(head *ListNode) *ListNode {
	slow := head
	fast := head
	cy := false
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			cy = true
			break
		}
	}
	if cy {
		slow = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	} else {
		return nil
	}
}