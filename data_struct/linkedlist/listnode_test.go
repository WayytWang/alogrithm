package linkedlist

import (
	"testing"
)

func TestGenByArray(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				s: []int{1, 2, 3, 4},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenByArray(tt.args.s)
			t.Log(got)
		})
	}
}

func TestListNode_Add(t *testing.T) {
	l := GenByArray([]int{1})
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				index: 1,
				val:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l.Add(tt.args.index, tt.args.val)
			t.Log(l)
		})
	}
}
