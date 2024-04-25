package stack

import (
	"strconv"
)

// MinStack 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	content []int
	min     []int // 每次入栈时 如果数据是最小值就入min栈
}

func Constructor() MinStack {
	return MinStack{
		content: make([]int, 0),
		min:     make([]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	m.content = append(m.content, val)
	if len(m.min) == 0 {
		m.min = append(m.min, val)
		return
	}
	mV := m.min[len(m.min)-1]
	if val <= mV {
		m.min = append(m.min, val)
		return
	}
}

func (m *MinStack) Pop() {
	tV := m.Top()
	m.content = m.content[:len(m.content)-1]
	mV := m.min[len(m.min)-1]
	if tV == mV {
		m.min = m.min[:len(m.min)-1]
	}
}

func (m *MinStack) Top() int {
	return m.content[len(m.content)-1]
}

func (m *MinStack) GetMin() int {
	return m.min[len(m.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// + - * /
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		tokenInt, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, tokenInt)
			continue
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		switch token {
		case "+":
			stack = append(stack, b+a)
		case "-":
			stack = append(stack, b-a)
		case "*":
			stack = append(stack, b*a)
		case "/":
			stack = append(stack, b/a)
		}
	}
	return stack[len(stack)-1]
}

// 示例 1： 输入：s = "3[a]2[bc]" 输出："aaabcbc"
// 示例 2： 输入：s = "3[a2[c]]" 输出："accaccacc"
// 示例 3： 输入：s = "2[abc]3[cd]ef" 输出："abcabccdcdcdef"
// 示例 4： 输入：s = "abc3[cd]xyz" 输出："abccdcdcdxyz"
func decodeString(s string) string {
	result := ""
	stack := make([]byte, 0)
	startIndex := -1
	// 找括号内的内容
	for _, v := range s {
		if v == ']' {
			for i := len(stack) - 1; i > 0; i-- {
				// startIndex 到 i 为括号内的内容
				if stack[i] == '[' {
					startIndex = i
					break
				}
			}
			temp := string(stack[startIndex+1:])
			// [ 前一位或者多位一定是数字
			// nIndex表示[ 前第一个数字的索引
			nIndex := 0
			for i := startIndex - 2; i > 0; i-- {
				n := stack[i]
				// 往前找到不为数字的索引
				if !(n == '0' || n == '1' || n == '2' || n == '3' || n == '4' || n == '5' || n == '6' || n == '7' || n == '8' || n == '9') {
					nIndex = i + 1
					break
				}
			}
			times, _ := strconv.Atoi(string(stack[nIndex:startIndex]))
			newTemp := ""
			for i := 0; i < times; i++ {
				newTemp += temp
			}
			temp = newTemp
			stack = stack[:nIndex]
			stack = append(stack, []byte(temp)...)
			continue
		}
		stack = append(stack, byte(v))
	}
	result = string(stack)
	return result
}

type BNode struct {
	Val   int
	Left  *BNode
	Right *BNode
}

// GenBTByArray 使用数组生成二叉树 不使用递归
func GenBTByArray(a []int) *BNode {
	if len(a) == 0 {
		return nil
	}
	rootV := a[0]
	root := &BNode{Val: rootV}
	if len(a) == 1 {
		return root
	}
	for i := 1; i < len(a); i++ {
		v := a[i]
		target := root // 插入新数据的父节点
		for target != nil {
			if v >= target.Val {
				if target.Right != nil {
					target = target.Right
				} else {
					target.Right = &BNode{Val: v}
					target = nil
				}
			} else {
				if target.Left != nil {
					target = target.Left
				} else {
					target.Left = &BNode{Val: v}
					target = nil
				}
			}
		}
	}
	return root
}

func LevelOrder(node *BNode) [][]int {
	if node == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := []*BNode{node}
	for len(queue) > 0 {
		levelLength := len(queue)
		levelResult := make([]int, levelLength)
		for i := 0; i < levelLength; i++ {
			levelResult[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelLength:]
		result = append(result, levelResult)
	}
	return result
}

func PreOrder(node *BNode) []int {
	result := make([]int, 0)
	stack := []*BNode{node}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		result = append(result, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return result
}

func InOrder(node *BNode) []int {
	result := make([]int, 0)
	stack := []*BNode{node}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		left := top
		for left != nil {
			stack = append(stack, left.Left)
			left = left.Left
		}
		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return result
}

func PostOrder(node *BNode) []int {
	return nil
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	cloneNode := &Node{}
	//stack := make(*Node, 0)
	//curNode := node
	//for curNode != nil {
	//	newNode := &Node{
	//		Val: curNode.Val,
	//	}
	//	for i := len(curNode.Neighbors) - 1; i > 0; i-- {
	//
	//	}
	//	for i, neighbor := range curNode.Neighbors {
	//
	//	}
	//}
	return cloneNode
}
