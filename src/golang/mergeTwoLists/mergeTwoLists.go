package mergeTwoLists

// 题目：https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// 解法
// 比较链表头节点，选小的作为新的链表头节点，并更新原链表头节点为Next
// 遍历链表比较，更新新链表cur节点
// 直到一个链表为空，把另一个链表所有节点更新到新链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var newListNode, cur *ListNode

	if l1.Val < l2.Val {
		newListNode = l1
		cur = l1
		l1 = l1.Next
	} else {
		newListNode = l2
		cur = l2
		l2 = l2.Next
	}

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
	}

	if l2 != nil {
		cur.Next = l2
	}

	return newListNode
}

// times: 1
