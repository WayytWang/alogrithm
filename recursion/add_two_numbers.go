package recursion

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := transferInt(l1, 0) + transferInt(l2, 0)
	return transferList(result, &ListNode{})
}

func transferInt(l *ListNode, index int) int {
	if l == nil {
		return 0
	}
	base := int(math.Pow10(index))
	return base*l.Val + transferInt(l.Next, index+1)
}

func transferList(number int, l *ListNode) *ListNode {
	// 本次节点
	node := l
	for  {
		// 计算本节点的值
		ele := number % 10
		node.Val = ele
		// 下一次迭代
		if number <= 0 {
			break
		}
		node.Next = &ListNode{}
		node = node.Next

	}
	return l
}
