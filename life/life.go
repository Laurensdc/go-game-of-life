package life

func countNeighbors(state [][]bool, x, y int) int {
	type Direction struct {
		x int
		y int
	}
	directions := []Direction{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	neighborCount := 0

	for _, d := range directions {
		if x+d.x < 0 {
			continue
		}
		if x+d.x >= len(state) {
			continue
		}
		if y+d.y < 0 {
			continue
		}
		if y+d.y >= len(state[0]) {
			continue
		}

		if state[x+d.x][y+d.y] == true {
			//	fmt.Printf("Found %v - %v %v %v\n", state[x+d.x][y+d.y], x, y, d)
			neighborCount++
		}
	}

	return neighborCount
}

func Tick(state [][]bool) [][]bool {
	// Copy over state to response
	// we don't want to mess with neighbors
	// while looping over them by mutating
	response := make([][]bool, len(state))
	for i := range state {
		response[i] = append([]bool(nil), state[i]...)
	}

	for i := range state {
		for j := range state[i] {
			// fmt.Printf("%v %v %v %v\n", state[i][j], i, j, len(state[i]))
			if countNeighbors(state, i, j) < 2 {
				response[i][j] = false
			}
			if countNeighbors(state, i, j) > 3 {
				response[i][j] = false
			}
			if state[i][j] == false && countNeighbors(state, i, j) == 3 {
				response[i][j] = true
			}

		}
	}

	return response
}
