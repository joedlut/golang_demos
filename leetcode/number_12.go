package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	h := head
	length := 0
	for h != nil {
		length++
		h = h.Next
	}
	if length == 1 && n == 1 {
		head = nil
		return head
	}
	if length == n {
		head = head.Next
		return head
	}

	deletePos := length - n
	tempPos := head
	for i := 1; i < deletePos; i++ {
		tempPos = tempPos.Next
	}
	tempPos.Next = tempPos.Next.Next
	return head
}

func main() {
}
