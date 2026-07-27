func plusOne(digits []int) []int {
    digit := ""

	// mengubah slice menjadi value menggunakan string (digit = 1234)
	for i := range digits {
		a := strconv.Itoa(digits[i])
		digit = digit+a
	}

	// mengubah string menjadi int agar bisa di tambah 1 
	num, _ := strconv.Atoi(digit)
	// 1235/1000
	c := num+1

	// butuh panjang c buat looping, jadi kita ubah dulu ke string buat dapetin itu
	d := strconv.Itoa(c)
	b := len(d)

	// fmt.Println(d)
	// fmt.Println(b)

	// karena udah dapat len nya maka kita ubah lagi jadi int
	// num2, _ := strconv.Atoi(d)
	
	output := make([]string, b)
	output2 := []int{}

	for i := 0; i < b; i++{
		output[i] = string(d[i])
	}

	for i := 0; i< b; i++ {
		z,_ := strconv.Atoi(output[i])
		output2 = append(output2, z)
	}

	fmt.Println(output2)


	return output2
}
