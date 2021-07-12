package sort

import (
	"reflect"
	"testing"
)

func TestSelectASC(t *testing.T) {
	type args struct {
		origin []Comparable
	}
	tests := []struct {
		name string
		args args
		want []Comparable
	}{
		{
			name: "选择排序升序结果测试",
			args: args{
				origin: []Comparable{CInt(6),CInt(5),CInt(4),CInt(3),CInt(2),CInt(1)},
			},
			want: []Comparable{CInt(1),CInt(2),CInt(3),CInt(4),CInt(5),CInt(6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectASC(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectASC() = %v, want %v", got, tt.want)
			}
		})
	}
}
