package linkedlist

import (
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 1, 2, 2, 3, 4, 4}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 1, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteDuplicates(tt.args.head)
			t.Log(tt.args.head)
		})
	}
}

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 1, 2, 2, 3, 4, 4}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 1, 1}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 2, 2, 3}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := deleteDuplicatesII(tt.args.head)
			t.Log(head)
		})
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenByArray([]int{6, 5, 4, 3, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			t.Log(got)
		})
	}
}

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenByArray([]int{6, 5, 4, 3, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList2(tt.args.head)
			t.Log(got)
		})
	}
}

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head:  GenByArray([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head:  GenByArray([]int{1, 2, 3, 4, 5}),
				left:  1,
				right: 5,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head:  GenByArray([]int{3, 5}),
				left:  1,
				right: 2,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head:  GenByArray([]int{1, 2, 3}),
				left:  3,
				right: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.args.head, tt.args.left, tt.args.right)
			t.Log(got)
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				list1: GenByArray([]int{1, 2, 4}),
				list2: GenByArray([]int{1, 3, 4}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			t.Log(got)
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.head, tt.args.x)
			t.Log(got)
		})
	}
}

func Test_findMiddle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 2, 3, 4, 5}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 2, 3, 4, 5, 6}),
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: GenByArray([]int{1, 2}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMiddle(tt.args.head)
			t.Log(got)
		})
	}
}

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				GenByArray([]int{6, 5, 4, 3, 2, 1}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortList(tt.args.head)
			t.Log(got)
		})
	}
}

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				GenByArray([]int{1, 2, 3, 4, 5}),
			},
		},
		{
			name: "",
			args: args{
				GenByArray([]int{1, 2, 3, 4}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			t.Log(tt.args.head)
		})
	}
}
