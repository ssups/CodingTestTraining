import (
    "sort"
    "math"
)

func solution(a int, b int, c int, d int) int {
    hashMap := make(map[int]int);
    keys := make([]int, 0);
    
    for _, each := range []int{a, b, c, d}{
        if _, ok:= hashMap[each]; ok {
            hashMap[each]++;
        } else {
            hashMap[each] = 1;
            keys = append(keys, each);
        }
    }
    
    switch len(keys) {
    case 1:
        return 1111 * keys[0];
    case 2:
        if hashMap[keys[0]] == 2 {
            return (keys[0] + keys[1]) * int(math.Abs(float64(keys[0] - keys[1])));
        }
        var (
            p = keys[0]
            q = keys[1]
        ) 
        if hashMap[keys[0]] != 3 {
            p = keys[1];
            q = keys[0];
        }
        return (10 * p + q) * (10 * p + q);
    case 3:
        for i, k := range keys {
            if hashMap[k] == 2 {
                 keys[i] = keys[len(keys) - 1];
                keys = keys[:len(keys) - 1];
                break
            }
        }
        return keys[0] * keys[1];
    case 4:
        sort.Ints(keys);
        return keys[0];
    default: 
        panic("wrong")
    }
}