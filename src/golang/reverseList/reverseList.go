package reverseList

// 题目：https://leetcode-cn.com/problems/reverse-linked-list/
// 反转一个单链表。

// 解法
// 分解原链表，每次逆序将节点增加到新的链表头，原链表头更新为Next

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

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	return newHead
}

// times : 1
