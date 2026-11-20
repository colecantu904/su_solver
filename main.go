package main

import (
	"fmt"

	sudoku_solver "github.com/colecantu904/su_solver/sudoko_solver"
)

func main() {
	puzzles := sudoku_solver.Getdata(2)

	fmt.Println("Puzzles output")
	fmt.Println(puzzles)

	fmt.Println("\nPuzzle lens")
	fmt.Println(len(puzzles[0].Puzzle), len(puzzles[0].Solution))

	fmt.Println("\nSolve output")
	fmt.Println(sudoku_solver.Solve(puzzles[0].Puzzle))
	fmt.Println("\nhere we will solve sudoku!")
}
