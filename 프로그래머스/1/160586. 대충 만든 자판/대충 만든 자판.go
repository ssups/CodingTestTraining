import (
    "strings"
)

func solution(keymap []string, targets []string) []int {
    shortCutKeyMap := make(map[string]int);
    for _, each := range keymap {
        for i, key := range strings.Split(each, "") {
            if v, ok := shortCutKeyMap[key]; ok && v <= i + 1 {
                continue;
            }                 
            shortCutKeyMap[key] = i + 1;
        } 
    }
    
    result := make([]int, len(targets));
    for i, each := range targets {
        for _, targetKey := range strings.Split(each, "") {
            v, ok := shortCutKeyMap[targetKey];
            if !ok {
                result[i] = -1;
                break
            }
            result[i] += v;  
        }
    }
    
    return result;
}