package sort

// 升序排列

// 冒泡排序
func bubble(ori []int) []int {
	for times := 0; times < len(ori); times++ {
		swap := false
		for i := 0; i < len(ori)-times-1; i++ {
			if ori[i] > ori[i+1] {
				ori[i], ori[i+1] = ori[i+1], ori[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return ori
}

// 插入排序
func insertion(ori []int) []int {
	for seg := 0; seg < len(ori)-1; seg++ {
		v := ori[seg+1] // 未排区第一个元素
		insert := 0
		for i := seg; i >= 0; i-- {
			if v < ori[i] {
				// i元素移动一位
				ori[i+1] = ori[i]
			} else {
				insert = i + 1
				break
			}
		}
		ori[insert] = v
	}
	return ori
}

// 选择排序
func selection(ori []int) []int {
	for seg := 0; seg < len(ori)-1; seg++ {
		min := ori[seg]
		selectIndex := seg
		for i := seg + 1; i < len(ori); i++ {
			if ori[i] < min {
				min = ori[i]
				selectIndex = i
			}
		}
		ori[seg], ori[selectIndex] = ori[selectIndex], ori[seg]
	}
	return ori
}

// 归并排序
func merge(ori []int) []int {
	if len(ori) == 1 {
		return ori
	}
	middle := len(ori) / 2
	p1 := merge(ori[:middle])
	p2 := merge(ori[middle:])
	// 合并
	result := make([]int, len(ori))
	i, j := 0, 0
	compare := 0
	for i < len(p1) && j < len(p2) {
		v1 := p1[i]
		v2 := p2[j]
		if v1 > v2 {
			result[compare] = v2
			j++
		} else {
			result[compare] = v1
			i++
		}
		compare++
	}
	// p1有剩余
	if i < len(p1) {
		for _, v := range p1[i:] {
			result[compare] = v
			compare++
		}

	}
	if j < len(p2) {
		for _, v := range p2[j:] {
			result[compare] = v
			compare++
		}
	}
	return result
}

// 快速排序
func quick(ori []int) []int {
	if len(ori) == 1 {
		return ori
	}
	seg := partition(ori)
	if seg > 0 {
		quick(ori[:seg])
	}
	if seg < len(ori)-1 {
		quick(ori[seg+1:])
	}
	return ori
}

// partition 最后一个元素作为分隔点 比它小的放左边 比它大的放右边 用尽量少的空间
func partition(ori []int) int {
	pivot := ori[len(ori)-1]
	// i表示已排区与未排区的分界线
	// j表示从未排区挑选元素 用于与pivot比较 决定是否要加入已排区
	i, j := 0, 0
	for j < len(ori)-1 {
		v := ori[j]
		if v < pivot {
			ori[i], ori[j] = ori[j], ori[i]
			i++
		}
		j++
	}
	ori[i], ori[len(ori)-1] = ori[len(ori)-1], ori[i]
	return i
}
