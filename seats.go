func assignSet() {
	file, err := os.Open("./seats.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inputs := make([]string, 0)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	max := 0
	all := make([]int, 0)

	for _, input := range inputs {
		arr := make([]int, 0)
		for i := 0; i < 128; i++ {
			arr = append(arr, i)
		}
		col := binarySearch(arr, input[0:7])[0]
		row := binarySearch([]int{0, 1, 2, 3, 4, 5, 6, 7}, input[7:])[0]
		fmt.Println(col*8 + row)
		if col*8+row > max {
			max = col*8 + row
		}
		all = append(all, col*8+row)
	}
	fmt.Println(max)
	sort.Ints(all)
	fmt.Println(all)
	i := 11
	for _, j := range all {
		if i != j {
			fmt.Println(i)
		}
		i++
	}
}

func binarySearch(input []int, str string) []int {
	fmt.Println(str)
	for _, run := range str {
		if run == 'F' || run == 'L' {
			input = input[0 : len(input)/2]

		}
		if run == 'B' || run == 'R' {
			input = input[len(input)/2 : len(input)]
		}
		fmt.Println(input)
	}

	fmt.Println(input)
	return input
}
