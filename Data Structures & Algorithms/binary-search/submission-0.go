func search(nums []int, target int) int {

	exists := -1

	for i := range nums {
		if target == nums[i] {
			exists = i
		}
	}

	return exists
}