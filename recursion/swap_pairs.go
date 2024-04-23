package recursion

import "algorithm/data_struct/linkedlist"

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 1->2->3->4  ==> 2->1->4->3
func swapPairs(head *linkedlist.ListNode) *linkedlist.ListNode {
	cur := head
	pre := &linkedlist.ListNode{
		Val:  0,
		Next: head,
	}
	headSwap := false
	for cur != nil {
		// cur 1
		next := cur.Next // next 2
		if next != nil {
			if !headSwap {
				head = next
			}
			// cur 与 next交换位置 pre.Next = next cur.Next = next.Next next.Next = cur
			pre.Next = next      // 0 -> 2 -> 3 ->4
			cur.Next = next.Next // 1 -> 3 ->4
			next.Next = cur
			headSwap = true
		}
		// 0 - 2 - 1 - 3 - 4  cur = 1
		pre = cur
		cur = cur.Next
	}
	return head
}

// 1->2->3->4
// 1 swap(2->3->4)
//	2->3->4
//	2 swap(3->4)
//		3->4
//		3 swap(4)
//			return 4
//      4->3 (1-2-4-3)
//		return 4

func swapPairsRec(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {

	}
	return head
}

func swapPairs1(head *linkedlist.ListNode) *linkedlist.ListNode {
	// 只有一个节点
	if head == nil || head.Next == nil {
		return head
	}
	// 只有两个节点，交换顺序
	if head.Next.Next == nil {
		// 交换head 与 head.Next顺序
		mn := head.Next
		head.Next.Next = head
		head.Next = nil
		return mn
	}
	// 交换前两个
	mn := head.Next
	head.Next = swapPairs1(head.Next.Next)
	mn.Next = head
	return mn
}
