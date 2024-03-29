package sort

type Comparable interface {
	Equal(comparable Comparable) bool
	MT(comparable Comparable) bool
	LT(comparable Comparable) bool

	// 小于等于
	LE(comparable Comparable) bool
	// 大于等于
	ME(comparable Comparable) bool
}

type CInt int

func (c CInt) Equal(comparable Comparable) bool {
	com,ok := comparable.(CInt)
	if !ok {
		panic("invalid type")
	}
	return int(c) == int(com)
}

func (c CInt) MT(comparable Comparable) bool {
	com,ok := comparable.(CInt)
	if !ok {
		panic("invalid type")
	}
	return int(c) > int(com)
}

func (c CInt) LT(comparable Comparable) bool {
	com,ok := comparable.(CInt)
	if !ok {
		panic("invalid type")
	}
	return int(c) < int(com)
}

func (c CInt) LE(comparable Comparable) bool {
	com,ok := comparable.(CInt)
	if !ok {
		panic("invalid type")
	}
	return int(c) <= int(com)
}

func (c CInt) ME(comparable Comparable) bool {
	com,ok := comparable.(CInt)
	if !ok {
		panic("invalid type")
	}
	return int(c) >= int(com)
}



