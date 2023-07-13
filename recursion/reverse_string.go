package recursion

// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
// https://leetcode.cn/problems/reverse-string

func reverseString(s []string) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		targetIndex := length - 1 - i
		s[i], s[targetIndex] = s[targetIndex], s[i]
	}
	return
}

// 递归方式
func reverseString2(s []string) {
	if len(s) == 1 {
		return
	}
	left := s[1:]
	s = append(s[1:], s[0])
	reverseString2(left)
}

func r(s []string) {

}
