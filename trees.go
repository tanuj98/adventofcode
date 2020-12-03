func trees(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := make([][]rune, 0)
	for scanner.Scan() {
		var rn []rune
		for x := 0; x < 11*8; x++ {
			rn = append(rn, []rune(scanner.Text())...)
		}
		res = append(res, rn)
	}

	fmt.Println(recurseTrees(0, 0, res))
}

func recurseTrees(x int, y int, arr [][]rune) int {
	if x >= len(arr) {
		return 0
	}

	count := 0
	if arr[x][y] == '#' {
		count++
	}

	return recurseTrees(x+2, y+1, arr) + count
}
