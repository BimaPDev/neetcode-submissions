func longestConsecutive(nums []int) int {
	// first you build the set
	// second you find x-1

	numsHash := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		numsHash[nums[i]] = struct{}{}
	}
	best := 0
	for x := range numsHash {
		if _,ok := numsHash[x-1]; !ok {
			length := 1
			for {
				if _, ok2 := numsHash[x+length]; !ok2 {
					break
				}
				length++
			}
			if length > best {
				best = length
			}
		}
	}
	return best
}