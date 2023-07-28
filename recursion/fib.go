package recursion

// 斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。
// F(0)=1，F(1)=1, F(n)=F(n - 1)+F(n - 2)（n ≥ 2，n ∈ N*）
// https://leetcode.cn/problems/fibonacci-number/

func fib(n int) int {
	cache := make(map[int]int)
	return f(n, cache)
}

func f(n int, cache map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	result, ok := cache[n]
	if ok {
		return result
	}
	r := f(n-1, cache) + f(n-2, cache)
	cache[n] = r
	return r
}
