func solution(n int) int64 {


    if n == 1{
        return 1
    }else if n == 2 {
        return 2
    }
    trial := make([]int, n)
    trial[0] = 1
    trial[1] = 2
    for i := 0 ; i < n -2 ;i++{
        trial[i + 2] = (trial[i+1] + trial[i]) %1234567
    }

    return int64(trial[n-1])
}