package sort

// 插入排序

// 数据交换
func InsertASCSwap(origin []Comparable) []Comparable {
	n := len(origin)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0; j-- {
			if origin[j].LT(origin[j-1]) {
				origin[j], origin[j-1] = origin[j-1], origin[j]
			}
		}
	}
	return origin
}

// 数据移动
// 待排序序列在长度为1000的情况下做性能测试，移动比交换要快不少
func InsertASCMove(origin []Comparable) []Comparable {
	n := len(origin)
	for i := 0; i < n-1; i++ {
		// this:待排区间第一个元素
		this := origin[i+1]
		j := i + 1
		for ; j > 0; j-- {
			if this.LT(origin[j-1]) {
				// move
				origin[j] = origin[j-1]
			} else {
				// this元素无需继续比较
				break
			}
		}
		origin[j] = this
	}
	return origin
}
