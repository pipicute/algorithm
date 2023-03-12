package largest_rectangle_in_histogram

import "container/list"

func largestRectangleArea(heights []int) int {
	stack := list.New()
	r := 0

	heights = append(heights, 0) // 右边界
	for i, v := range heights {
		for stack.Len() != 0 {
			topIndex := stack.Back().Value.(int)
			if heights[topIndex] < v {
				break
			}
			if heights[topIndex] == v {
				stack.Remove(stack.Back())
				break
			}
			stack.Remove(stack.Back())
			preTopIndex := -1 // 左边界
			if stack.Len() != 0 {
				preTopIndex = stack.Back().Value.(int)
			}
			r = max(r, (i-preTopIndex-1)*heights[topIndex])
		}
		stack.PushBack(i)
	}
	return r
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 思路：对每一个元素，分别求左右侧的边界（比当前值小），算出最大面积
// 构建单调栈（从小到大），当遍历的当前元素小于栈顶 i 时，i 的左右边界便可确定
