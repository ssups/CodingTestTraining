import "sort"

func solution(k int, tangerine []int) int {
    sizeList := []int{}
    amountMap := make(map[int]int)
    
    for _, size := range tangerine {
        if _, ok := amountMap[size]; !ok {
            sizeList = append(sizeList, size)
        } 
        amountMap[size]++
    }
    
    sort.Slice(sizeList, func(i, j int) bool { 
        return amountMap[sizeList[i]] > amountMap[sizeList[j]] 
    })
    
    i := 0
    for k > 0 {
        k = k - amountMap[sizeList[i]]
        i ++
    }
    
    return i
}