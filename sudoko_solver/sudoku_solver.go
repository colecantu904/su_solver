package sudoku_solver

// so what is the approach here? what is the proper method?
// with this solution checker, we are just assuming the 3x3 method
func Solve(pr string) string {
	// ps for possibilities, this will be the slice to store possible answers
	var ps [9][9]map[int]bool

	// ps[0] = map[int]bool{1: true}
	var p Puzzle = RawToPuzzle(pr)

	for i := 0; i < len(p); i = i + 1 {
		for j := 0; j < len(p); j = j + 1 {
			// if the space is already filled skip
			if p[i][j] != 0 {
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

		}

	}

	return PuzzleToRaw(p)
}
