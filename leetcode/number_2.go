package main

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。


示例 1：


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
 

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode
	var head1 *ListNode
	var head2 *ListNode
	i := 1
	var add [101]byte
	for head1, head2 = l1, l2; head1 != nil && head2 != nil; i++ {
		var newNode *ListNode = &ListNode{
			Val:  0,
			Next: nil,
		}
		if add[i-1] == 1 {
			if head1.Val+head2.Val+1 > 9 {
				newNode.Val = head1.Val + head2.Val + 1 - 10
				add[i] = 1
			} else {
				newNode.Val = head1.Val + head2.Val + 1
			}
		} else {
			if head1.Val+head2.Val > 9 {
				newNode.Val = head1.Val + head2.Val - 10
				add[i] = 1
			} else {
				newNode.Val = head1.Val + head2.Val
			}
		}
		newNode.Next = nil
		if head == nil {
			head = newNode
			cur = head
		} else {
			cur.Next = newNode
			cur = cur.Next
		}
		//fmt.Println(add)
		head1 = head1.Next
		head2 = head2.Next
	}
	if head1 != nil {
		for ; head1 != nil; i++ {
			var val int
			if add[i-1] == 1 {
				if head1.Val+1 > 9 {
					val = head1.Val + 1 - 10
					//fmt.Println(i)
					add[i] = 1
				} else {
					val = head1.Val + 1
				}
				var newNode *ListNode = &ListNode{
					Val:  val,
					Next: nil,
				}
				cur.Next = newNode
				cur = cur.Next
			} else {
				cur.Next = head1
				break
			}
			head1 = head1.Next
		}
	}

	if head2 != nil {
		for ; head2 != nil; i++ {
			var val int
			if add[i-1] == 1 {
				if head2.Val+1 > 9 {
					val = head2.Val + 1 - 10
					add[i] = 1
				} else {
					val = head2.Val + 1
				}
				var newNode *ListNode = &ListNode{
					Val:  val,
					Next: nil,
				}
				cur.Next = newNode
				cur = cur.Next
			} else {
				cur.Next = head2
				break
			}
			head2 = head2.Next
		}
	}

	//最后一位一定要记得进位
	if add[i-1] == 1 {
		var newNode *ListNode = &ListNode{
			Val:  1,
			Next: nil,
		}
		cur.Next = newNode
		cur = cur.Next

	}

	return head
}
