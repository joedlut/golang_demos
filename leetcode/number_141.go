package main

/*func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	m := make(map[*ListNode]int)
	p := head
	m[p]++
	if p.Next == nil {
		return false
	}
	p = p.Next
	for p != nil {
		m[p]++
		if m[p] == 2 {
			return true
		}
		p = p.Next
	}
	return false
}
*/

/*
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return false
}
*/

// 龟兔赛跑算法  定义一个慢指针 一次移动一步，一个快指针 一次移动两步

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		//一个结点 或者两个结点的情况
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
