package main

import "fmt"

type A struct {
	Name string
}

// 指针再次赋值
func pointer() {
	a1 := &A{
		Name: "wang",
	}
	fmt.Println(a1)
	a2 := a1
	fmt.Println(a2)
	a2.Name = "wyj"
	fmt.Println(a1)
	fmt.Println(a2)
	a2 = &A{
		Name: "wang2",
	}
	fmt.Println(a2)
	fmt.Println(a1)
}

// 切片作为传值
func sliceParam(s []string) {
	top := s[0]
	left := s[1:]
	left[0] = "*"
	s = append(left, top)
	fmt.Println(s)
}

func main() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
	//s := []string{"1", "2", "3", "4"}
	//fmt.Println(s)
	//sliceParam(s)
	//fmt.Println(s)
	//
	//sliceParam(s[2:])
	//fmt.Println(s)
}
