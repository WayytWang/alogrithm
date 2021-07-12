package sort

// 冒泡排序

func BubbleASC(origin []Comparable) []Comparable {
	n := len(origin)
	for i := 0; i < n; i++ {
		changed := false
		for j := n - 1; j > 0; j-- {
			if origin[j].LT(origin[j-1]) {
				origin[j],origin[j-1] = origin[j-1],origin[j]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return origin
}

func BubbleDES(origin []Comparable) []Comparable {
	n := len(origin)
	for i := 0; i < n; i++ {
		changed := false
		for j := n - 1; j > 0; j-- {
			if origin[j].MT(origin[j-1]) {
				origin[j],origin[j-1] = origin[j-1],origin[j]
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return origin
}