package sort

func SelectASC(origin []Comparable) []Comparable {
	n := len(origin)
	for i := 0; i < n-1; i++ {
		// 待排区间第一个元素
		index := i
		min := origin[index]
		for j := i + 1; j < n; j++ {
			if origin[j].LT(min) {
				min = origin[j]
				index = j
			}
		}
		origin[i], origin[index] = origin[index], origin[i]
	}
	return origin
}
