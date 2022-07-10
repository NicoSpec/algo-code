package readed

import "algo-code/readed"

// 最大盛水容器
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	res := 0

	for i < j {
		if height[i] < height[j] {
			res = readed.max(res, (j - i) * height[i])
			i = i + 1
		} else {
			res = readed.max(res, (j - i) * height[j])
			j = j - 1
		}
	}
	return res
}

func max2(pre, next int) int {
	if pre < next {
		return next
	}
	return pre
}