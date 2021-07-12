package sort

import (
	"reflect"
	"testing"
)


func getComparableList(length int) []Comparable {
	list := make([]Comparable,0)
	for i := length; i > 0; i-- {
		list = append(list, CInt(i))
	}
	return list
}

// 2462137
func BenchmarkInsertASCSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertASCSwap(getComparableList(1000))
	}
}

// 1958274
func BenchmarkInsertASCMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertASCMove(getComparableList(1000))
	}
}

func TestInsertASCSwap(t *testing.T) {
	type args struct {
		origin []Comparable
	}
	tests := []struct {
		name string
		args args
		want []Comparable
	}{
		{
			name: "插入排序升序结果测试（数据交换）",
			args: args{
				origin: []Comparable{CInt(6),CInt(5),CInt(4),CInt(3),CInt(2),CInt(1)},
			},
			want: []Comparable{CInt(1),CInt(2),CInt(3),CInt(4),CInt(5),CInt(6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertASCSwap(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertASCSwap() = %v, want %v", got, tt.want)
			}else{
				t.Log(got)
			}
		})
	}
}

func TestInsertASCMove(t *testing.T) {
	type args struct {
		origin []Comparable
	}
	tests := []struct {
		name string
		args args
		want []Comparable
	}{
		{
			name: "插入排序升序结果测试（数据移动）",
			args: args{
				origin: []Comparable{CInt(6),CInt(5),CInt(4),CInt(3),CInt(2),CInt(1)},
			},
			want: []Comparable{CInt(1),CInt(2),CInt(3),CInt(4),CInt(5),CInt(6)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertASCMove(tt.args.origin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertASCMove() = %v, want %v", got, tt.want)
			}else{
				t.Log(got)
			}
		})
	}
}