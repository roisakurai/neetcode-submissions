func getConcatenation(nums []int) []int {
	ans:=[]int{}
	
	for i := range nums {
		ans = append(ans, nums[i])		
	}

	merged := append(nums, ans...)

	return merged
}
