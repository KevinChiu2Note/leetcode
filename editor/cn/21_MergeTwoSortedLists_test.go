package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
// 示例： 
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
// 
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归法-1
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeTwoSortedLists(t *testing.T) {
	node11 := &ListNode{Val: 1}
	node12 := &ListNode{Val: 2}
	node13 := &ListNode{Val: 4}
	node11.Next = node12
	node12.Next = node13

	node21 := &ListNode{Val: 1}
	node22 := &ListNode{Val: 3}
	node23 := &ListNode{Val: 4}
	node21.Next = node22
	node22.Next = node23

	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, mergeTwoLists(node11, node21).getAllListNodeToArray())
}

// 迭代法-1
func mergeTwoLists_1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
				cur = cur.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
				cur = cur.Next
			}
			continue
		}
		if l1 == nil && l2 != nil {
			cur.Next = l2
			break
		}
		if l2 == nil && l1 != nil {
			cur.Next = l1
			break
		}
		if l1 == nil && l2 == nil {
			break
		}
	}
	return dummy.Next
}

// 迭代法-2
func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
