func hasDuplicate(nums []int) bool {
     dupes := make(map[int]struct{}) // make an empty hash

     for _, end := range nums {
        _, exists := dupes[end]
        if exists {
            return true
        }
        dupes[end] = struct {}{}
     }
     return false
}
