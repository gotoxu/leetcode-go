package leetcode

func lengthOfLIS(nums []int) int {
	return lengthOfLISOfN(nums, len(nums))
}

func lengthOfLISOfN(nums []int, i int) int {
	a := 1
	for j := 0; j < i; j++ {
		if nums[j] < nums[i] {
			a = max(a, lengthOfLISOfN(nums, j)+1)
		}
	}
	return a
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
