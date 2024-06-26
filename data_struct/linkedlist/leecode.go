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

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为： L0 → L1 → … → Ln - 1 → Ln
// 请将其重新排列后变为： L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func reorderList(head *ListNode) {
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil

	var node *ListNode
	// 翻转tail列表
	for tail != nil {
		node = &ListNode{
			Val:  tail.Val,
			Next: node,
		}
		tail = tail.Next
	}

	headCur := head
	tailCur := node

	for headCur != nil && tailCur != nil {
		tailNext := tailCur.Next
		headNext := headCur.Next
		// tailCur插入headCur之后
		tailCur.Next = headCur.Next
		headCur.Next = tailCur
		tailCur = tailNext
		headCur = headNext
	}
	return
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			slow = head
			fast = fast.Next
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表 。如果是，返回 true ；否则，返回 false 。
func isPalindrome(head *ListNode) bool {
	var reverseNode *ListNode
	cur := head
	for cur != nil {
		reverseNode = &ListNode{
			Val:  cur.Val,
			Next: reverseNode,
		}
		cur = cur.Next
	}
	for head != nil && reverseNode != nil {
		if head.Val != reverseNode.Val {
			return false
		}
		head = head.Next
		reverseNode = reverseNode.Next
	}
	return true
}

func copyList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for head != nil {
		pre.Next = &ListNode{
			Val: head.Val,
		}
		pre = pre.Next
		head = head.Next
	}
	return dummy.Next
}

//给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
//构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。
//新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
//例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。
//返回复制链表的头节点。
//用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
//
//val：一个表示 Node.val 的整数。
//random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
//你的代码 只 接受原链表的头节点 head 作为传入参数。

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mMap := make(map[*Node]int)
	iMap := make(map[int]*Node)
	index := 1
	copyDummy := &Node{}
	pre := copyDummy
	cur := head
	for cur != nil {
		mMap[cur] = index
		pre.Next = &Node{
			Val: cur.Val,
		}
		iMap[index] = pre.Next
		pre = pre.Next
		cur = cur.Next
		index++
	}
	cur = head
	cCur := copyDummy.Next
	for cur != nil {
		i := mMap[cur.Random]
		rn := iMap[i]
		cCur.Random = rn
		cur = cur.Next
		cCur = cCur.Next
	}
	return copyDummy.Next
}
