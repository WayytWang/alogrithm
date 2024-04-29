package binary_search

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
		},
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.args.nums, tt.args.target)
			t.Log(got)
		})
	}
}
