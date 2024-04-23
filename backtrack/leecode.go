package backtrack

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]

func subsets(nums []int) [][]int {
	result := [][]int{[]int{}}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

		}
	}
	return result
}
