func maxArea(heights []int) int {
	length := len(heights)
	left := 0
	right := length - 1
	best := 0

	for left < right {
		area := min(heights[left], heights[right]) * (right - left)
		if area > best {
			best = area 
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return best
}