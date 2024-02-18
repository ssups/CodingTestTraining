import "math"

func solution(n int, a int, b int) int {    
    if a > b {
        temp := a
        a = b
        b = temp
    }
    
    count := 0
    
    for {
        count ++
        if b - a == 1 && a % 2 == 1 {
            return count
        }
        
        b = int(math.Ceil(float64(b) / float64(2)))
        a = int(math.Ceil(float64(a) / float64(2)))
    }
    
    panic("wrong")
}   