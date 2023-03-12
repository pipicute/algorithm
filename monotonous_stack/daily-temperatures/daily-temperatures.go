package daily_temperatures

import "container/list"

func dailyTemperatures(temperatures []int) []int {
	s := list.New()
	r := make([]int, len(temperatures))
	for i, v := range temperatures {
		for s.Len() > 0 {
			topIndex := s.Back().Value.(int)
			if temperatures[topIndex] >= v {
				break
			}
			s.Remove(s.Back())
			r[topIndex] = i - topIndex
		}
		s.PushBack(i)
	}
	return r
}

// https://leetcode.com/problems/daily-temperatures
// 对数组每个元素，求出下一个更大元素的下标(差值)
