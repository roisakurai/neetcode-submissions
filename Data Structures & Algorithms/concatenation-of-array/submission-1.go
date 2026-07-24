func getConcatenation(nums []int) []int {
	ans:=[]int{}
	
	for i := range nums {
		ans = append(ans, nums[i])		
	}

	ans = append(ans, nums...)

	return ans
}
