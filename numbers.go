unc numbers() {
	// input := []int{0, 3, 6}
	//14,8,16,0,1,17
	m := make(map[int][]int)
	// m[3] = []int{1}
	// m[1] = []int{2}
	// m[2] = []int{3}
	m[14] = []int{1}
	m[8] = []int{2}
	m[16] = []int{3}
	m[0] = []int{4}
	m[1] = []int{5}
	m[17] = []int{6}
	turn := 7
	last := 17
	for ; turn <= 30000000; turn++ {
		if l, exist := m[last]; exist && len(l) >= 2 {
			last = l[len(l)-1] - l[len(l)-2]
		} else {
			if len(l) == 0 {
				m[last] = make([]int, 0)
			}

			last = 0
		}

		m[last] = append(m[last], turn)
	}

	fmt.Println(last)
}
