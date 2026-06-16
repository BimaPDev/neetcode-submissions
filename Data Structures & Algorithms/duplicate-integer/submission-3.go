func hasDuplicate(nums []int) bool {
    dupes := make(map[int]struct{})
    for _, num := range nums {
        dupes[num] = struct{}{}
    }
    return len(dupes) < len(nums)
}
