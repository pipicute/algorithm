package sum_of_distances

func distance(nums []int) []int64 {
	groups := make(map[int][]int, 0)
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := make([]int, n+1) // 前缀和， s[0] = 0
		for i, x := range a {
			s[i+1] = s[i] + x // 前缀和，x 为下标
		}
		for i, target := range a {
			left := target*i - s[i]             // 左侧面积（将距离之和转换为二维柱状图集合，target*i 表示左侧元素都从 0 开始算距离和）
			right := s[n] - s[i] - target*(n-i) // 绿色面积
			ans[target] = int64(left + right)
		}
	}
	return ans
}

// https://leetcode.com/contest/weekly-contest-340/problems/sum-of-distances/
