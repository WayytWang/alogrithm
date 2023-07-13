package sort

// merge_sort(p...r) = merge(merge_sort(p...q),merge_sort(q+1...r))

func MergeASC(origin []Comparable) []Comparable {
	length := len(origin)
	return mergeASC(origin, 0, length-1)
}

func mergeASC(origin []Comparable, s, e int) []Comparable {
	if e <= s {
		return []Comparable{origin[s]}
	}
	middle := (s + e) / 2
	return merge(mergeASC(origin, s, middle), mergeASC(origin, middle+1, e))
}

func merge(l1 []Comparable, l2 []Comparable) []Comparable {
	i := 0
	j := 0
	merged := make([]Comparable, len(l1)+len(l2))
	for index := 0; index < len(merged); index++ {
		var v1 Comparable
		var v2 Comparable
		if i < len(l1) {
			v1 = l1[i]
			v2 = v1
		}
		if j < len(l2) {
			v2 = l2[j]
		}
		if v1.LE(v2) {
			merged[index] = v1
			i++
		} else {
			merged[index] = v2
			j++
		}
	}
	return merged
}
