import (
    "strconv"
)

func solution(n_str string) int {
    res, _ := strconv.ParseInt(n_str, 10, 32)
    return int(res)
}