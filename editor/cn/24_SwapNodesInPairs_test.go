package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
//
// 
//
// 示例: 
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
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
// 递归法 - 优化
func swapPairs(head *ListNode) *ListNode {
	// 特殊情况，空链或者只有一个节点的链
	if head == nil || head.Next == nil {
		return head
	}

	result := head.Next
	head.Next = swapPairs(result.Next)
	result.Next = head

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSwapNodesInPairs(t *testing.T) {
	link := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}
	l := swapPairs(link)
	assert.Equal(t, l.getAllListNodeToArray(), []int{1, 2, 7, 9})
}

// 迭代法
func swapPairs_1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	// head的前一个节点
	prev := dummy
	for head != nil && head.Next != nil {
		firstNode := head
		secondNode := head.Next
		// 交换
		prev.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		// 移动下标
		head = firstNode.Next
		prev = firstNode
	}
	return dummy.Next
}

// 递归法
func swapPairs_2(head *ListNode) *ListNode {
	// 特殊情况，空链或者只有一个节点的链
	if head == nil || head.Next == nil {
		return head
	}

	// 要交换的节点
	firstNode := head
	secondNode := head.Next

	// 交换
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}
