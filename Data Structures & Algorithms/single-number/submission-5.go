func singleNumber(nums []int) int {
    count := make(map[int]int)

    for _, v := range nums {
        count[v]++
    }

    for i, v := range count {
        if v == 1 {
            return i
        }
    }

    return 0
}
