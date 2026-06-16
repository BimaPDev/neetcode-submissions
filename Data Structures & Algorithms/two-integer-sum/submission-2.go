func twoSum(nums []int, target int) []int {
	numHash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		need := target - nums[i]

		if j, ok := numHash[need]; ok {
			return []int{j, i}
		}

		numHash[nums[i]] = i
	}

	return nil
}
