package biTree

import (
	"testing"
)

func TestNode_Add(t *testing.T) {
	n := NewNode(66)
	n.Add(11)
	n.Add(55)
	n.Add(33)
	n.Add(99)
	n.Add(44)
	n.Add(77)
	n.Add(22)
	n.Add(88)
	n.Add(9)
	n.Add(10)

	t.Log(n.LevelOrderFor())
	//t.Log(n.InOrder())
	//t.Log(n.InOrderFor())
	//t.Log(n.PostOrder())
	//t.Log(n.PostOrderFor())
	//t.Log(n.PreOrder())
	//t.Log(n.PreOrderFor())
}
