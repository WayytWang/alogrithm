package stack

import (
	"strconv"
)

type MinStack struct {
	content []int
}

func Constructor() MinStack {
	return MinStack{
		content: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.content = append(this.content, val)
}

func (this *MinStack) Pop() {
	this.content = this.content[:len(this.content)-1]
}

func (this *MinStack) Top() int {
	return this.content[len(this.content)-1]
}

func (this *MinStack) GetMin() int {
	min := this.content[0]
	for i, v := range this.content {
		if v < min {
			min = this.content[i]
		}
	}
	return min
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

// 示例 1：
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"

// 示例 2：
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"

// 示例 3：
// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"

// 示例 4：
// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"
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
