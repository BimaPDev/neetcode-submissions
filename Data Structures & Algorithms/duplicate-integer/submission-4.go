func hasDuplicate(nums []int) bool {
    dupes := make(map[int]struct{})
    for _, i := range nums {
        dupes[i] = struct {}{}
    }
    return len(dupes) < len(nums)
}
