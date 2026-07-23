func singleNumber(nums []int) int {
    seen := make(map[int]int)
    
    for _, num := range nums {
        seen[num]++
    }

    for a := range seen {
        if seen[a] == 1 {
            return a
        }
    }

    return 0
}
