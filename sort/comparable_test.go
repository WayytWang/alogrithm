package sort

import "testing"

func TestCInt_Equal(t *testing.T) {
	type args struct {
		comparable Comparable
	}
	tests := []struct {
		name string
		c    CInt
		args args
		want bool
	}{
		{
			name: "true",
			c:    5,
			args: args{
				comparable: CInt(5),
			},
			want: true,
		},
		{
			name: "false",
			c:    5,
			args: args{
				comparable: CInt(6),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Equal(tt.args.comparable); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}


