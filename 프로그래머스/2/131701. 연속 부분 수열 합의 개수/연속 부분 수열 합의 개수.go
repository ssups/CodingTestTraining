func solution(elements []int) int {
    numsMap := make(map[int]struct{})
    
    for i := 1; i <= len(elements); i++ {
        target := append(elements, elements[:i-1]...)
        for j := 0; j < len(elements); j++ {
            numsMap[sum(target[j:j+i])] = struct{}{}    
        }
    }
    
    return len(numsMap)
}

func sum(elements []int) int {
    res := 0
    for _, el := range elements {
        res+= el
    }
    
    return res
}