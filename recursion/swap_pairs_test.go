package recursion

import (
	"algorithm/data_struct/linkedlist"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *linkedlist.ListNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode
	}{
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1, 2}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1, 2, 3}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1, 2, 3, 4}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.args.head)
			t.Log(got)
		})
	}
}

func Test_swapPairs2(t *testing.T) {
	type args struct {
		head *linkedlist.ListNode
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.ListNode
	}{
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1, 2}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1, 2, 3}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: linkedlist.GenByArray([]int{1, 2, 3, 4}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs2(tt.args.head)
			t.Log(got)
		})
	}
}
