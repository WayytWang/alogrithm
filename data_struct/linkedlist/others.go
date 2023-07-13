package linkedlist

import "fmt"

// GenByArray 从数组生成链表
func GenByArray(s []int) *ListNode {
	dummyNode := &ListNode{}
	pre := dummyNode
	for i := range s {
		pre.Next = &ListNode{
			Val: s[i],
		}
		pre = pre.Next
	}
	return dummyNode.Next
}

// String 链表输出成字符串
func (l *ListNode) String() string {
	if l.Next == nil {
		return fmt.Sprintf("%d", l.Val)
	}
	return fmt.Sprintf("%d->%s", l.Val, l.Next.String())
}
