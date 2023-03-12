package trapping_rain_water

import (
	"algorithm/utils"
	"container/list"
)

func trap1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	r := 0
	leftHeight, rightHeight := utils.GetMaxNumList(height)
	for i, v := range height {
		minHeight := utils.Min(leftHeight[i], rightHeight[i])
		if minHeight > v {
			r += minHeight - v
		}
	}
	return r
}

func trap2(height []int) int {
	r := 0
	stack := list.New()
	for i, v := range height {
		for stack.Len() > 0 {
			topIndex := stack.Back().Value.(int)
			if height[topIndex] > v {
				break
			}
			stack.Remove(stack.Back())
			if height[topIndex] == v {
				continue
			}
			if stack.Len() > 0 {
				preTopIndex := stack.Back().Value.(int)
				r += (i - preTopIndex - 1) * (utils.Min(v, height[preTopIndex]) - height[topIndex]) // 只求中间宽度，需要减一！
			}
		}
		stack.PushBack(i)
	}
	return r
}

// https://leetcode.com/problems/trapping-rain-water/
// 元素下标为 x，元素值为 y，宽度为 1 组成柱形图，求雨水可以填满的区域面积
// 思路：①双指针(纵向处理) ②单调栈(横向处理)
// 纵向处理：每一列的高度为左右两侧峰值的较小高度
// 横向处理：算凹槽 ——> 栈中递减(存索引)，遇到较大值计算宽高，其中高为差值，宽为中间宽度
