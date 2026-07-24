func search(nums []int, target int) int {
	output := 0
	for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				output = i
				return output
		}
	}

	return -1
}
