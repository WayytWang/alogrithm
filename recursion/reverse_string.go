package recursion

func reverseString(s []string) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func reverseStringRec(s []string, res *[]string) {
}
