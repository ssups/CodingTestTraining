import "sort"

type process struct {
    Location int
    Priority int
}

func solution(priorities []int, location int) int {
    sorted := make([]int, len(priorities))
    copy(sorted, priorities)
    sort.Slice(sorted, func(i, j int) bool { return sorted[i] > sorted[j] })
    
    
    queue := make([]process, len(priorities))
    for i, priority := range priorities {
        queue[i] = process{Location: i, Priority: priority}
    }
    
    count := 0
    for {
        runProcess := indexOf(queue, sorted[0])
        count ++
        if queue[runProcess].Location == location {
            return count
        }
        
        sorted = sorted[1:]
        queue = append(queue[runProcess+1:], queue[:runProcess]...)
    }
    
    panic("fail")
}

func indexOf(queue []process, target int) int {    
    for i, process := range queue {
        if process.Priority == target {
            return i
        }
    }
    
    panic("not included target")
}

