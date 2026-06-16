func productExceptSelf(nums []int) []int {
	// go forward
	forwardArray := make([]int, len(nums))
	if len(nums) > 0{
		forwardArray[0] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		forwardArray[i] = forwardArray[i-1] * nums[i]
	}
	//fmt.Println(forwardArray)

	// backwardArray
	backwardArray := make([]int, len(nums))
	if len(nums) > 0 {
		backwardArray[len(nums)-1] = nums[len(nums)-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		backwardArray[i] = backwardArray[i+1] * nums[i]
	}
	//fmt.Println(backwardArray)

	// Calculation
	result := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			result[i] = backwardArray[i+1]
		}else if i == len(nums)-1 {
			result[i] = forwardArray[i-1]
		} else {
			result[i] = forwardArray[i-1] * backwardArray[i+1]
		}
	}
	//fmt.Println(result)

	return result
}