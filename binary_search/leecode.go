package binary_search

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//
//示例 1: 输入: nums = [-1,0,3,5,9,12], target = 9 输出: 4
//示例 2: 输入: nums = [-1,0,3,5,9,12], target = 2 输出: -1

// 你可以假设 nums 中的所有元素是不重复的。 n 将在 [1, 10000]之间。 nums 的每个元素都将在 [-9999, 9999]之间。
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (end - start) / 2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			mid = mid + (end-mid)/2
			start = mid
		} else {
			mid = (mid - start) / 2
			end = mid
		}
	}
	return -1
}
