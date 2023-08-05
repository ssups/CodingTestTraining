func solution(num_list []int) int {
    for i, v := range num_list {
        if v < 0 {return i}
    }
    return -1
}