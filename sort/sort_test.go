package sort

import (
	"reflect"
	"testing"
)

func Test_bubble(t *testing.T) {
	type args struct {
		ori []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				ori: []int{4, 5, 6, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bubble(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertion(t *testing.T) {
	type args struct {
		ori []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				ori: []int{4, 5, 3, 6, 1, 2},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{2, 4, 6, 3},
			},
			want: []int{2, 3, 4, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertion(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selection(t *testing.T) {
	type args struct {
		ori []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				ori: []int{4, 5, 3, 6, 1, 2},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{2, 4, 6, 3},
			},
			want: []int{2, 3, 4, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selection(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test(t *testing.T) {
	ori := []int{1, 2}
	t.Log(len(ori) / 2)
	t.Log(ori[:0])
}

func Test_merge(t *testing.T) {
	type args struct {
		ori []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				ori: []int{11, 8, 3, 9, 7, 1, 2, 5},
			},
			want: []int{1, 2, 3, 5, 7, 8, 9, 11},
		},
		{
			name: "",
			args: args{
				ori: []int{11, 8, 3, 9, 7, 1, 2},
			},
			want: []int{1, 2, 3, 7, 8, 9, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		ori []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				ori: []int{1, 2, 3, 4, 5, 6},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				ori: []int{6, 5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				ori: []int{6, 11, 3, 9, 8},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
			t.Log(tt.args.ori)
		})
	}
}

func Test_quick(t *testing.T) {
	type args struct {
		ori []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				ori: []int{4, 5, 3, 6, 1, 2},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{2, 4, 6, 3},
			},
			want: []int{2, 3, 4, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			args: args{
				ori: []int{6, 11, 3, 9, 8},
			},
			want: []int{3, 6, 8, 9, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quick(tt.args.ori); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quick() = %v, want %v", got, tt.want)
			}
		})
	}
}
