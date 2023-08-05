func solution(arr []int, k int) []int {
    newArr := []int{}
    tempM := make(map[int]struct{})
    
    for _, v := range arr {      
        if _, ok := tempM[v]; !ok {
            tempM[v] = struct{}{}
            newArr = append(newArr, v)
        }
        if len(newArr) >= k {break}
    }
    
    gap := k - len(newArr)
    if gap > 0 {
        for i := 0; i < gap; i++ {
            newArr = append(newArr, -1)
        }
    }
    
    return newArr
}