package recursion

import (
	"testing"
)

func Test_transferInt(t *testing.T) {
	type args struct {
		l     *ListNode
		index int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				l:     &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				index: 0,
			},
			want: 342,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transferInt(tt.args.l, tt.args.index); got != tt.want {
				t.Errorf("transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transferList(t *testing.T) {
	type args struct {
		number int
		l      *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				number: 1234,
				l:      &ListNode{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 got := transferList(tt.args.number, tt.args.l)
			 for got.Next != nil {
			 	t.Log(got.Val)
			 }
		})
	}
}
