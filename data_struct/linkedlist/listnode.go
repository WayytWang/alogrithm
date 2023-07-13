package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Add 在index位置插入节点(从0开始)
func (l *ListNode) Add(index int, val int) {
	node := l
	// 找到index的前一个元素
	for i := 0; i < index; i++ {
		if node == nil {
			panic(fmt.Sprintf("不够%d个元素", index+1))
		}
		node = node.Next
	}
	node.Next = &ListNode{
		Val:  val,
		Next: node.Next,
	}
}
