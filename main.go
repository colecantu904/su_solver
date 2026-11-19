package main

import (
	"fmt"

	"github.com/colecantu904/su_solver/sudoko_solver"
)

func main() {
	puzzles := sudoku_solver.Getdata(2)
	fmt.Println(len(puzzles[0].Puzzle), len(puzzles[0].Solution), len(puzzles[1].Puzzle), len(puzzles[1].Solution))
	fmt.Println("here we will solve sudoku!")
}
