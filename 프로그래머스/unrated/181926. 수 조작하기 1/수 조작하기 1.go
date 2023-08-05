import (
    "strings"
)

func solution(n int, control string) int {
    controls := strings.Split(control,"")
    for _, v := range controls {
        switch v {
            case "w":
                n += 1
            case "s":
                n -= 1
            case "d":
                n += 10
            case "a":
                n -= 10
        }
    }
    return n
}