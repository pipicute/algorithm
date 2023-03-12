package count_the_number_of_beautiful_subarrays

func beautifulSubArrays(nums []int) int64 {
	pre := []int{0}
	for i, v := range nums {
		pre = append(pre, pre[i]^v)
	}
	r := int64(0)
	t := make(map[int]int64)
	for _, v := range pre {
		r += t[v]
		t[v]++
	}
	return r
}

// https://leetcode.com/contest/weekly-contest-336/problems/count-the-number-of-beautiful-subarrays/
// 子数组满足：二进制位值'1'为偶数个 ——> 异或和为 0 的子数组数目 ——> 前缀和异或为 0 的子数组数目 ——> 前缀和元素相同的数对
