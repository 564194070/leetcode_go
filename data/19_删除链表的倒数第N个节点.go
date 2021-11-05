package data

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//注意面对链表问题的时候 头节点可能被删除 记得新增头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headbefore := &ListNode{0, head}
	left, right := headbefore, headbefore

	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return headbefore.Next
}
