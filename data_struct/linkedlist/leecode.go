package linkedlist

// 给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。 题目数据保证链表已经按升序 排列
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		if cur.Val == next.Val {
			// 删除 next 节点
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -10000,
		Next: head,
	}
	cur := dummy
	delFlag := false
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		if next != nil && cur.Val == next.Val {
			// 删除 next 节点
			cur.Next = next.Next
			delFlag = true
		} else {
			// 如果当前Val是被删除过的 需要删除cur节点
			if delFlag {
				pre.Next = cur.Next
				delFlag = false
				cur = cur.Next
			} else {
				pre = cur
				cur = cur.Next
			}
		}
	}
	return dummy.Next
}

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	stack := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	dummy := &ListNode{}
	pre := dummy
	for i := len(stack) - 1; i >= 0; i-- {
		pre.Next = &ListNode{
			Val: stack[i].Val,
		}
		pre = pre.Next
	}
	return dummy.Next
}

// O(1)空间复杂度
func reverseList2(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		cur = &ListNode{
			Val:  head.Val,
			Next: cur,
		}
		head = head.Next
	}
	return cur
}

// 递归实现
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sh := reverseList3(head.Next)
	// 将head移动至sh的尾部
	var tail = sh
	for tail.Next != nil {
		tail = tail.Next
	}
	head.Next = nil
	tail.Next = head
	return sh
}

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	index := 1
	cur := head
	var reverseCur *ListNode
	var startNode *ListNode
	var endNode *ListNode
	for cur != nil {
		if index == left-1 {
			startNode = cur
		}
		if index >= left && index <= right {
			node := &ListNode{
				Val:  cur.Val,
				Next: reverseCur,
			}
			if index == left {
				// endNode是反转链的最后一个节点
				endNode = node
			}
			reverseCur = node
			if index == right {
				// reverseCur为一条反转链
				if startNode != nil {
					startNode.Next = reverseCur
				} else {
					head = reverseCur
				}
				if endNode != nil {
					endNode.Next = cur.Next
				}
				break
			}
		}
		cur = cur.Next
		index++
	}
	return head
}

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	cur1 := list1
	cur2 := list2
	for cur1 != nil && cur2 != nil {
		v1 := cur1.Val
		v2 := cur2.Val
		node := &ListNode{}
		if v1 < v2 {
			// 取list1上的值
			node.Val = v1
			cur1 = cur1.Next
		} else {
			node.Val = v2
			cur2 = cur2.Next
		}
		pre.Next = node
		pre = pre.Next
	}
	if cur1 != nil {
		pre.Next = cur1
	}
	if cur2 != nil {
		pre.Next = cur2
	}
	return dummy.Next
}

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。 你应当 保留 两个分区中每个节点的初始相对位置。
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{
		Val:  -1000,
		Next: head,
	}
	pre := dummy
	cur := head

	nDummy := &ListNode{}
	nPre := nDummy
	for cur != nil {
		if cur.Val >= x {
			// 将cur移动至一条新的链表
			pre.Next = cur.Next
			nPre.Next = &ListNode{
				Val: cur.Val,
			}
			nPre = nPre.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	pre.Next = nDummy.Next
	return dummy.Next
}

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMiddle(head)
	tail := mid.Next
	mid.Next = nil
	left := sortList(head)
	right := sortList(tail)

	cur1 := left
	cur2 := right

	nDummy := &ListNode{}
	pre := nDummy
	for cur1 != nil && cur2 != nil {
		v1 := cur1.Val
		v2 := cur2.Val
		if v1 < v2 {
			pre.Next = &ListNode{
				Val: v1,
			}
			cur1 = cur1.Next
		} else {
			pre.Next = &ListNode{
				Val: v2,
			}
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	if cur1 != nil {
		pre.Next = cur1
	}
	if cur2 != nil {
		pre.Next = cur2
	}
	return nDummy.Next
}

// 找到中点后 需要将链表切割成两端  head:middle middle:final
// 为了切割 需要将middle的Next置空 head:middle包含了middle
// middle:last 这后一段 应该表示为 middle.Next:final (middle.Next为置空前的值)
// 当整个链表长度为2时 如果取出middle为第二个值 在切割时 前一段 1->2 第二段为空 在递归切割中 会陷入循环递归
// 所以在长度为2时 middle应该取第一个值  在切割时 前一段 1->nil 后一段为2->nil
// 得出结论 偶数长度时 middle取靠前值 如 1-2-3-4-5-6 中点取3 不取4
func findMiddle(head *ListNode) *ListNode {
	fast := head.Next
	middle := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		middle = middle.Next
	}
	return middle
}
