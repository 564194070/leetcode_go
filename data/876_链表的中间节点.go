package data

/* Definition for singly-linked list.*/
//双指针快慢指针处理链表中间节点问题
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	left := head
	right := head

	for right.Next != nil {

		right = right.Next
		if right.Next != nil {
			left = left.Next
			right = right.Next
		} else {
			left = left.Next
		}

	}
	return left
}
