package next_greater_element_i

import "container/list"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	s := list.New()
	t := make(map[int]int)
	for _, v := range nums2 {
		for s.Len() > 0 {
			topV := s.Back().Value.(int)
			if v <= topV {
				break
			}
			s.Remove(s.Back())
			t[topV] = v
		}
		s.PushBack(v)
	}
	r := make([]int, len(nums1))
	for i, v := range nums1 {
		if nextV, exist := t[v]; exist {
			r[i] = nextV
		} else {
			r[i] = -1
		}
	}
	return r
}

// https://leetcode.com/problems/next-greater-element-i/
// 对nums1每个元素，求出nums2中下一个更大元素的值(所有元素值互不相同)
// 思路：①求出数组2每个元素的下一个更大元素 ——> 单调栈； ②遍历数组1，哈希查找数组2中的下一个更大元素
