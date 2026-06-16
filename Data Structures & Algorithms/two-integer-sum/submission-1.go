func twoSum(nums []int, target int) []int {
    prevMap := map[int]int{}

    for i, n := range nums {
        diff := target - n
        if j, ok := prevMap[diff]; ok {
             return []int{j, i}
        }
        prevMap[n] = i
    }
    return nil
}
