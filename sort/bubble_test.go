package sort

import (
	"reflect"
	"testing"
)

func BenchmarkBubbleASC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleASC([]Comparable{CInt(1),CInt(2),CInt(3),CInt(4),CInt(5)})
	}
}

func TestBubbleASC(t *testing.T) {
	type args struct {
		origin []Comparable
	}
	tests := []struct {
		name string
		args args
		want []Comparable
	}{
		{
			name: "冒泡排序升序结果测试",
			args: args{
				origin: []Comparable{CInt(2),CInt(3),CInt(1),CInt(9),CInt(7)},
			},
			want: []Comparable{CInt(1),CInt(2),CInt(3),CInt(7),CInt(9)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleASC(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}else {
				t.Log(got)
			}
		})
	}
}

func TestBubbleDES(t *testing.T) {
	type args struct {
		origin []Comparable
	}
	tests := []struct {
		name string
		args args
		want []Comparable
	}{
		{
			name: "冒泡排序降序结果测试",
			args: args{
				origin: []Comparable{CInt(2),CInt(3),CInt(1),CInt(9),CInt(7)},
			},
			want: []Comparable{CInt(9),CInt(7),CInt(3),CInt(2),CInt(1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleDES(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}else {
				t.Log(got)
			}
		})
	}
}