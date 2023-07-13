package recursion

import (
	"algorithm/data_struct/linkedlist"
)

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
// 1->2->3->4  ==> 2->1->4->3
// https://leetcode.cn/problems/swap-nodes-in-pairs/

func swapPairs(head *linkedlist.ListNode) *linkedlist.ListNode {
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
	head.Next = swapPairs(head.Next.Next)
	mn.Next = head
	return mn
}

func swapPairs2(head *linkedlist.ListNode) *linkedlist.ListNode {
	// 只有一个节点
	if head == nil || head.Next == nil {
		return head
	}
	mn := swapPairs2(head.Next)
	// 交换head与mn的位置
	head.Next = mn.Next
	mn.Next = head
	return mn
}
