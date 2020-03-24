package cn

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。 
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 
//
// 
//
// 示例： 
//
// 给你这个链表：1->2->3->4->5 
//
// 当 k = 2 时，应当返回: 2->1->4->3->5 
//
// 当 k = 3 时，应当返回: 3->2->1->4->5 
//
// 
//
// 说明： 
//
// 
// 你的算法只能使用常数的额外空间。 
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。 
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
// 递归法法
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	count := 0
	for cur != nil && count != k {
		cur = cur.Next
		count++
	}
	if count == k {
		cur = reverseKGroup(cur, k)
		for count != 0 {
			count--
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
		}
		head = cur
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseNodesInKGroup(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next, node2.Next, node3.Next, node4.Next = node2, node3, node4, node5
	assert.Equal(t, []int{2, 1, 4, 3, 5}, reverseKGroup(node1, 2).getAllListNodeToArray())
	assert.Equal(t, (*ListNode)(nil), reverseKGroup(nil, 2))
	node1 = &ListNode{Val: 1}
	node2 = &ListNode{Val: 2}
	node3 = &ListNode{Val: 3}
	node4 = &ListNode{Val: 4}
	node5 = &ListNode{Val: 5}
	node1.Next, node2.Next, node3.Next, node4.Next = node2, node3, node4, node5
	assert.Equal(t, []int{1, 2, 3, 4, 5}, reverseKGroup(node1, 6).getAllListNodeToArray())
	node1 = &ListNode{Val: 1}
	node2 = &ListNode{Val: 2}
	node3 = &ListNode{Val: 3}
	node4 = &ListNode{Val: 4}
	node5 = &ListNode{Val: 5}
	node1.Next, node2.Next, node3.Next, node4.Next = node2, node3, node4, node5
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reverseKGroup(node1, 5).getAllListNodeToArray())
}

// 使用栈
func reverseKGroup_1(head *ListNode, k int) *ListNode {
	stack := list.New()
	dummyNode := &ListNode{Next: head}
	writeIndex := dummyNode
	for {
		count := 0
		index := head
		for count < k && index != nil {
			stack.PushBack(index)
			index = index.Next
			count++
		}
		if count != k {
			break
		}
		for stack.Len() > 0 {
			if e := stack.Back(); e != nil {
				writeIndex.Next = e.Value.(*ListNode)
				writeIndex = writeIndex.Next
				stack.Remove(e)
			}
		}
		writeIndex.Next = index
		head = index

	}
	return dummyNode.Next
}

// 尾插法
func reverseKGroup_2(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre, tail := dummyNode, dummyNode

	for {
		count := 0
		for count != k && tail != nil {
			count++
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		head = pre.Next
		for pre.Next != tail {
			cur := pre.Next
			pre.Next = cur.Next
			cur.Next = tail.Next
			tail.Next = cur
		}
		pre = head
		tail = head
	}

	return dummyNode.Next
}
