package sudoku_solver

import (
	"log"
	"encoding/json"
	"os"
	"math/rand"
)

type Puzzle [81]int

type PuzzleExample struct {
	Puzzle Puzzle `json:"puzzle"`
	Solution Puzzle `json:"solution"`
	Difficulty string `json:"difficulty"`
}

// pass -1 for all, or pass n for at most n random puzzles
func Getdata( amount int ) []PuzzleExample {
	var puzzle_examples []PuzzleExample

	data, err := os.ReadFile("testdata/sudoku.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &puzzle_examples)
	if err != nil {
		log.Fatal(err)
	}

	if amount != -1 {
		rand.Shuffle(len(puzzle_examples), func( i, j int ) {
			puzzle_examples[i], puzzle_examples[j] = puzzle_examples[j], puzzle_examples[i]
		})
		puzzle_examples = puzzle_examples[0:amount]
	}

	return puzzle_examples
}
