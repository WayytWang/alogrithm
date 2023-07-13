package sort

import (
	"reflect"
	"testing"
)

func TestMergeASC(t *testing.T) {
	type args struct {
		origin []Comparable
	}
	tests := []struct {
		name string
		args args
		want []Comparable
	}{
		{
			name: "归并排序",
			args: args{
				origin: []Comparable{CInt(6), CInt(5), CInt(4), CInt(3), CInt(2), CInt(1)},
			},
			want: []Comparable{CInt(1), CInt(2), CInt(3), CInt(4), CInt(5), CInt(6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeASC(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeASC() = %v, want %v", got, tt.want)
			}
		})
	}
}
