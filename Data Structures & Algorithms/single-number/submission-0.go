func singleNumber(nums []int) int {
    seen := make(map[int]int)
    output := 0
    
    for _, num := range nums {
        seen[num]++
    }

    for a := range seen {
        if seen[a] == 1 {
            output = a
        }
    }

    return output
}
