package recursion

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				b: 0,
			},
			want: 0,
		},
		{
			name: "1",
			args: args{
				b: 1,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				b: 2,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				b: 3,
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				b: 4,
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				b: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.b); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
