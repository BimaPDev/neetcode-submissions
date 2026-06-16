func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS, countT := make(map[byte]int), make(map[byte]int)

    for i := 0; i < len(s); i++ {
        countS[s[i]]++
        countT[t[i]]++
    }

    for k, v := range countS{
        if countT[k] != v {
            return false
        }
    }

    return true
}
