package biTree

import (
	"testing"
)

func TestNode_Add(t *testing.T) {
	n := NewNode(1)
	n.Add(8)
	n.Add(6)
	n.Add(8)
	n.Add(0)
	n.Add(2)
	n.Add(4)
	n.Add(2)
	n.Add(6)
	n.Add(2)
	n.Add(1)

	t.Log(n.LevelOrderFor())
	t.Log(n.PreOrderFor())
	t.Log(n.LevelOrderFor())
	//t.Log(n.InOrder())
	//t.Log(n.InOrderFor())
	//t.Log(n.PostOrder())
	//t.Log(n.PostOrderFor())
	//t.Log(n.PreOrder())
	//t.Log(n.PreOrderFor())
}
