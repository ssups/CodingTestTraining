import "math/big"

func solution(n int) int64 {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    
    list := []*big.Int{big.NewInt(1), big.NewInt(2)}
    for i := 2; i < n; i ++ {
        list = append(list, new(big.Int).Add(list[i - 2], list[i - 1]))
    }
    
    return new(big.Int).Mod(list[len(list) - 1], big.NewInt(1234567)).Int64()
}