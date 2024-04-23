package recursion

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				s: []string{"h", "e", "l", "l", "o"},
			},
		},
		{
			name: "",
			args: args{
				s: []string{"b", "a", "c", "k", "u", "p"},
			},
		},
		{
			name: "",
			args: args{
				s: []string{"A", "*", "m", "a", "n", ",", "*", "a", "*", "p", "l", "a", "n", ",", "*", "a", "*", "c", "a", "n", "a", "l", ":", "*", "P", "a", "n", "a", "m", "a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			t.Log(tt.args.s)
		})
	}
}

//func Test_reverseString2(t *testing.T) {
//	type args struct {
//		s []string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{
//			name: "",
//			args: args{
//				s: []string{"w", "a"},
//			},
//		},
//		{
//			name: "",
//			args: args{
//				s: []string{"w", "a", "n"},
//			},
//		},
//		{
//			name: "",
//			args: args{
//				s: []string{"h", "e", "l", "l", "o"},
//			},
//		},
//		{
//			name: "",
//			args: args{
//				s: []string{"b", "a", "c", "k", "u", "p"},
//			},
//		},
//		{
//			name: "",
//			args: args{
//				s: []string{"A", "*", "m", "a", "n", ",", "*", "a", "*", "p", "l", "a", "n", ",", "*", "a", "*", "c", "a", "n", "a", "l", ":", "*", "P", "a", "n", "a", "m", "a"},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			reverseStringRec(tt.args.s)
//			t.Log(tt.args.s)
//		})
//	}
//}
