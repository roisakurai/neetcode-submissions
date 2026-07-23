func missingNumber(nums []int) int {
	sort.Ints(nums)
	output := 0
	for i, num := range nums {
		if i != num {
			output = i
			return output
		}

		output = i
	}

	return output+1

}
