package the_number_of_beautiful_subsets

func beautifulSubsets(nums []int, k int) int {
	r := -1
	t := make(map[int]int)

	var backTracking func(curIndex int)
	backTracking = func(curIndex int) {
		if curIndex == len(nums) {
			r++
			return
		}
		backTracking(curIndex + 1)
		if num := nums[curIndex]; t[num+k] == 0 && t[num-k] == 0 {
			t[num]++ // 可能有重复值，不能用 bool
			backTracking(curIndex + 1)
			t[num]--

		}
	}
	backTracking(0)
	return r
}

// https://leetcode.com/problems/the-number-of-beautiful-subsets/
