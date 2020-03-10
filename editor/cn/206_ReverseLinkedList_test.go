package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//反转一个单链表。
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) getAllListNodeToArray() []int {
	var res []int
	cur := node
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func TestReverseLinkedList(t *testing.T) {
	// 空链表测试
	assert.Equal(t, (*ListNode)(nil), reverseList(nil))
	// 只有一个节点测试
	single := &ListNode{Val: 1}
	assert.Equal(t, single, reverseList(single))
	// 多个
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
	assert.Equal(t, []int{7, 9, 1, 2}, reverseList(link).getAllListNodeToArray())
}

// 双指针法/迭代法
func reverseList_1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	next := head.Next
	cur.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	return cur
}

// 双指针法 - 优化
func reverseList_2(head *ListNode) *ListNode {
	// 指向前一个
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 双指针法 - 再优化
func reverseList_3(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		head, head.Next, pre = head.Next, pre, head
	}
	return pre
}

// 创建新的链表
func reverseList_4(head *ListNode) *ListNode {
	var list *ListNode
	for head != nil {
		list = &ListNode{
			Val:  head.Val,
			Next: list,
		}
		head = head.Next
	}
	return list
}
