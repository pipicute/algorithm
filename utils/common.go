package utils

import (
	"sort"
)

// GetMaxNumList 获取数组每个元素左右侧的最大值
// 核心思想：maxLeft[i] = max(height[i], maxLeft[i - 1]);
func GetMaxNumList[T int | int64](l []T) ([]T, []T) {
	left, right := make([]T, len(l)), make([]T, len(l))
	for i := 1; i < len(l); i++ {
		left[i] = Max(l[i-1], left[i-1])
	}
	for i := len(l) - 2; i >= 0; i-- {
		right[i] = Max(l[i+1], right[i+1])
	}
	return left, right
}

func Max[T int | int64](nums ...T) T {
	if len(nums) == 0 {
		panic(nums)
	}
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		if r < nums[i] {
			r = nums[i]
		}
	}
	return r
}

func Min[T int | int64](nums ...T) T {
	if len(nums) == 0 {
		panic(nums)
	}
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		if r > nums[i] {
			r = nums[i]
		}
	}
	return r
}

// GetPrimes 获取从 1 - maxNum 的质数（素数）
// 质数定义：只有1和它本身两个约数的数（1不是质数）
// 核心思想：标记
func GetPrimes[T int | int64](maxNum T) []T {
	r := make([]T, 0)
	np := make([]bool, maxNum+1)
	for i := T(2); i <= maxNum; i++ {
		if np[i] {
			continue
		}
		r = append(r, i)
		for j := i * i; j <= maxNum; j += i { // 下一个非质数由 i * i 开始，因为 (i-k) * i 已经在前面的循环中被标记
			np[j] = true
		}
	}
	return r
}

// IsPrime 判断一个数是否是质数
// 算法：若 x 不能被 2 ~ 根号x（下取整）整除，则 x 为质数
// -- 反证法：在上述条件下，若 x 能被 y（大于根号x）整除，那 x 也能被 (x/y) 整除，而 (x/y) < 根号x，矛盾。
func IsPrime[T int | int64](num T) bool {
	for i := T(2); i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1 // 1 不是质数
}

// SortIndex 返回从排序后的原数组索引
// 采用稳定排序，保证值相同的元素不乱序
func SortIndex[T int | int64](l []T, asc bool) []T {
	t := make([][2]T, len(l))
	for i := range l {
		t[i][0] = T(i)
		t[i][1] = l[i]
	}
	sort.SliceStable(t, func(i, j int) bool {
		if asc {
			return t[i][1] < t[j][1]
		}
		return t[i][1] > t[j][1]
	})
	r := make([]T, 0)
	for _, v := range t {
		r = append(r, v[0])
	}
	return r
}
