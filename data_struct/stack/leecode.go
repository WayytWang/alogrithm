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

// InOrder 中序遍历
// 先左，再自己，最后右节点 遍历到左节点后，左节点还有左右节点 深度搜索问题
// 拿到根节点 最左的路径全部入栈
func InOrder(node *BNode) []int {
	result := make([]int, 0)
	stack := make([]*BNode, 0)
	cur := node
	for cur != nil || len(stack) > 0 {
		left := cur
		for left != nil {
			stack = append(stack, left)
			left = left.Left
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, top.Val)
		cur = top.Right
	}
	return result
}

func PostOrder(node *BNode) []int {
	result := make([]int, 0)
	cur := node
	stack := make([]*BNode, 0)
	var lastVisitedNode *BNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == lastVisitedNode {
			result = append(result, top.Val)
			stack = stack[:len(stack)-1]
			lastVisitedNode = top
		} else {
			cur = top.Right
		}
	}
	return result
}

type Node struct {
	Val       int
	Neighbors []*Node
}

// NewGraph [[2,4],[1,3],[2,4],[1,3]]
func NewGraph(nn [][]int) *Node {
	nodeMap := make(map[int]*Node)
	for i := range nn {
		node := &Node{
			Val: i + 1,
		}
		nodeMap[i+1] = node
	}
	for i, ns := range nn {
		node := nodeMap[i+1]
		for _, n := range ns {
			node.Neighbors = append(node.Neighbors, nodeMap[n])
		}

	}
	return nodeMap[1]
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// 原node对应的新node
	visited := map[*Node]*Node{node: {Val: node.Val}}
	queue := []*Node{node}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		for i := range curNode.Neighbors {
			newN := visited[curNode.Neighbors[i]]
			if newN == nil {
				queue = append(queue, curNode.Neighbors[i])
				newN = &Node{Val: curNode.Neighbors[i].Val}
				visited[curNode.Neighbors[i]] = newN
			}
			visited[curNode].Neighbors = append(visited[curNode].Neighbors, newN)
		}
	}
	return visited[node]
}

func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return cloneRec(node, visited)
}

func cloneRec(node *Node, visited map[*Node]*Node) *Node {
	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for i := range node.Neighbors {
		newNe, ok := visited[node.Neighbors[i]]
		if !ok {
			newNe = cloneRec(node.Neighbors[i], visited)
		}
		newNode.Neighbors = append(newNode.Neighbors, newNe)
	}
	return newNode
}

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// 输入：grid = [
//
//	["1","1","1","1","0"],
//	["1","1","0","1","0"],
//	["1","1","0","0","0"],
//	["0","0","0","0","0"]
//
// ]
// 输出：1
// 输入：grid = [
//
//	["1","1","0","0","0"],
//	["1","1","0","0","0"],
//	["0","0","1","0","0"],
//	["0","0","0","1","1"]
//
// ]
// 输出：3
// 提示：
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
func numIslands(grid [][]byte) int {
	return 0
}
