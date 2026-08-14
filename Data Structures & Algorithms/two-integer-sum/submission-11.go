func twoSum(nums []int, target int) []int {
    list := make(map[int]int)

    for i, v := range nums {
        diff := target - v
        if j, found := list[diff]; found {
            return []int{j, i}
        }
        list[v] = i
    }
    return []int{}
}
