package stack

import (
	"reflect"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				tokens: []string{"2", "1", "+", "3", "*"},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
		{
			name: "",
			args: args{
				s: "abc3[cd]xyz",
			},
			want: "abccdcdcdxyz",
		},
		{
			name: "",
			args: args{
				s: "10[leetcode]",
			},
			want: "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_GetMin(t *testing.T) {
	//[[],[0],[1],[0],[],[],[]]
	ms := Constructor()
	ms.Push(0)
	ms.Push(1)
	ms.Push(0)
	t.Log(ms.GetMin())
	ms.Pop()
	t.Log(ms.GetMin())
}

func TestGenBTByArray(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want *BNode
	}{
		{
			name: "",
			args: args{
				a: []int{1, 8, 6, 8, 0, 2, 4, 2, 6, 2, 1},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenBTByArray(tt.args.a)
			t.Log(got)
		})
	}
}

func TestLevelOrder(t *testing.T) {
	type args struct {
		node *BNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				node: GenBTByArray([]int{1, 8, 6, 8, 0, 2, 4, 2, 6, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LevelOrder(tt.args.node)
			t.Log(got)
		})
	}
}

func TestPreOrder(t *testing.T) {
	type args struct {
		node *BNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				node: GenBTByArray([]int{1, 8, 6, 8, 0, 2, 4, 2, 6, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreOrder(tt.args.node)
			t.Log(got)
		})
	}
}

func TestInOrder(t *testing.T) {
	type args struct {
		node *BNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				node: GenBTByArray([]int{1, 8, 6, 8, 0, 2, 4, 2, 6, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InOrder(tt.args.node)
			t.Log(got)
		})
	}
}

func TestPostOrder(t *testing.T) {
	type args struct {
		node *BNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				node: GenBTByArray([]int{1, 8, 6, 8, 0, 2, 4, 2, 6, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PostOrder(tt.args.node)
			t.Log(got)
		})
	}
}

func TestNewGraph(t *testing.T) {
	type args struct {
		nn [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "",
			args: args{
				// [2,4],[1,3],[2,4],[1,3]
				nn: [][]int{{2, 4}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.nn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
