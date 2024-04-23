package biTree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (n *Node) Add(value int) (rn *Node) {
	if value >= n.Value {
		if n.Right == nil {
			rn = NewNode(value)
			n.Right = rn
		} else {
			rn = n.Right.Add(value)
		}
		return rn
	} else {
		if n.Left == nil {
			rn = NewNode(value)
			n.Left = rn
		} else {
			rn = n.Left.Add(value)
		}
		return rn
	}
}

func (n *Node) PreOrder() []int {
	r := make([]int, 0)
	r = append(r, n.Value)
	if n.Left != nil {
		r = append(r, n.Left.PreOrder()...)
	}
	if n.Right != nil {
		r = append(r, n.Right.PreOrder()...)
	}
	return r
}

func (n *Node) PreOrderFor() []int {
	r := make([]int, 0)
	stack := []*Node{n}
	for {
		if len(stack) == 0 {
			break
		}
		// pop
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, cn.Value)
		if cn.Right != nil {
			// 右节点push
			stack = append(stack, cn.Right)
		}
		if cn.Left != nil {
			// 左节点push
			stack = append(stack, cn.Left)
		}
	}
	return r
}

func (n *Node) InOrder() []int {
	r := make([]int, 0)
	if n.Left != nil {
		r = append(r, n.Left.InOrder()...)
	}
	r = append(r, n.Value)
	if n.Right != nil {
		r = append(r, n.Right.InOrder()...)
	}
	return r
}

func (n *Node) InOrderFor() []int {
	r := make([]int, 0)
	stack := make([]*Node, 0)
	cur := n
	// 最左的节点肯定是最先
	for cur != nil || len(stack) > 0 {
		// 保存这一支的最左节点
		leftNode := cur
		for leftNode != nil {
			stack = append(stack, leftNode)
			leftNode = leftNode.Left
		}
		// 目前的最左节点是第一个需要收集的节点
		topNode := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		r = append(r, topNode.Value)
		// 开始处理最左节点的右节点
		cur = topNode.Right
	}
	return r
}

func (n *Node) PostOrder() []int {
	r := make([]int, 0)
	if n.Left != nil {
		r = append(r, n.Left.PostOrder()...)
	}
	if n.Right != nil {
		r = append(r, n.Right.PostOrder()...)
	}
	r = append(r, n.Value)
	return r
}

func (n *Node) PostOrderFor() []int {
	r := make([]int, 0)
	stack := make([]*Node, 0)
	cur := n
	var lastNode *Node
	// 找到最左分支上无子节点的左节点 或者最左分支上 没有左兄弟的右节点
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		topNode := stack[len(stack)-1]
		// 右节点的判断
		// 右节点为空时表示topNode可以收集数据并且弹出了 弹出后 没有更新cur
		if topNode.Right == nil || topNode.Right == lastNode {
			stack = stack[0 : len(stack)-1]
			r = append(r, topNode.Value)
			lastNode = topNode
		} else {
			cur = topNode.Right
		}
	}
	fmt.Println(lastNode)
	return r
}

func (n *Node) LevelOrderFor() [][]int {
	r := make([][]int, 0)
	queue := []*Node{n}
	for len(queue) > 0 {
		level := make([]int, 0)
		nodeNum := len(queue)
		for i := 0; i < nodeNum; i++ {
			node := queue[i]
			level = append(level, node.Value)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if len(queue) == nodeNum {
			queue = make([]*Node, 0)
		} else {
			queue = queue[nodeNum:]
		}
		r = append(r, level)
	}
	return r
}

func isValidBST(root *Node) bool {
	if root == nil {
		return false
	}
	// 节点无左节点也无右节点
	if root.Right == nil && root.Left == nil {
		return true
	}
	// 节点无左节点有右节点
	if root.Left == nil && root.Right != nil {
		if root.Value > root.Right.Value {
			return false
		}
		return isValidBST(root.Right)
	}
	// 节点无右节点有左节点
	if root.Right == nil && root.Left != nil {
		if root.Value < root.Left.Value {
			return false
		}
		return isValidBST(root.Left)
	}
	// 节点既有左节点又有右节点
	if root.Right != nil && root.Left != nil {
		if root.Value > root.Right.Value {
			return false
		}
		if root.Value < root.Left.Value {
			return false
		}
		return isValidBST(root.Left) && isValidBST(root.Right)
	}
	return false
}
