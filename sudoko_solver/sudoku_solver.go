package sudoku_solver

func getBoxValues(idx int, p Puzzle) []int {
	return []int{1}
}

// so what is the approach here? what is the proper method?
// with this solution checker, we are just assuming the 3x3 method

// the agent made a good point, when we update a value into a space, we need to make sure to update
// all of the possiblities (ps) that are eliminated in the row, column, and box.

// considering this, we might need a different apporoach, instead lets just do one pass like this...

func Solve(pr string) string {
	// ps for possibilities, this will be the slice to store possible answers
	var ps [9][9]map[int]bool

	// ps[0] = map[int]bool{1: true}
	var p Puzzle = RawToPuzzle(pr)

	var solved bool = false
	var solved_counter int

	for !solved {
		for i := 0; i < len(p); i = i + 1 {
			for j := 0; j < len(p); j = j + 1 {
				// if the space is already filled skip, and increase the counter
				if p[i][j] != 0 {
					solved_counter++
					continue
				}

				// if possibilities have not been set then go ahead and make all numbes possible
				if ps[i][j] == nil {
					ps[i][j] = map[int]bool{
						1: true,
						2: true,
						3: true,
						4: true,
						5: true,
						6: true,
						7: true,
						8: true,
						9: true,
					}
				}

				// update the square if possible
				var possible int
				var found int
				for k := 1; k < 10; k++ {
					if ps[i][j][k] {
						possible++
						found = k
					}
				}

				// update the square and update any possibilities that are affected
				if possible == 1 {
					// update the square
					p[i][j] = found

					// update all possibilities affected
					for k := 0; k < len(ps[i]); k++ {
						if k == j {
							continue
						}

					}

				} else {
					// update all of the possibilites for the square

					continue

				}

			}

		}

		// if all spaces are filled clear as solved, reset the counter
		if solved_counter == 81 {
			solved = true
		} else {
			solved_counter = 0
		}
	}

	return PuzzleToRaw(p)
}
