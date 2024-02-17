import "fmt"

func solution(arr []int) int {
    lcm := arr[0]
    
    for i := 1; i < len(arr); i++ {
        fmt.Println(lcm)
        lcm = calcLCM(lcm, arr[i])
    }
    
    return lcm
}

func calcLCM(num1, num2 int) int {
    a, b := num1, num2
    if a < b {
        a, b = num2, num1
    }
    
    r := a % b
    for r > 0 {
        a = b
        b = r
        r = a % b
    }
    
    gcd := b
    return num1 * num2 / gcd
}