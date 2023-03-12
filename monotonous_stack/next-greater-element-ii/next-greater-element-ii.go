package next_greater_element_ii

import "container/list"

func nextGreaterElements(nums []int) []int {
	r := make([]int, len(nums))
	for i := range r {
		r[i] = -1
	}
	stack := list.New()
	for k := 0; k < 2; k++ {
		for i, v := range nums {
			for stack.Len() != 0 {
				topIndex := stack.Back().Value.(int)
				if nums[topIndex] >= v {
					break
				}
				stack.Remove(stack.Back())
				r[topIndex] = v
			}
			stack.PushBack(i)
		}
	}
	return r
}

// https://leetcode.com/problems/next-greater-element-ii/
// 对于循环数组 nums 的每个元素，获取比它大的下一个值，若不存在则用-1表示
// 思路：①最多循环一次，即可找到每个元素可能的下一个较大值； ②遍历两次两组，执行单调栈操作
