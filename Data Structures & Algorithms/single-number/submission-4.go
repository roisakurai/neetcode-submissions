func singleNumber(nums []int) int {
	seen := make(map[int]int)

	for _, i:=range nums {
		seen[i]++
	}

	fmt.Println(seen)

	for  q,a := range seen {
		if a == 1 {
			return q
		}
	}

return 0

}
