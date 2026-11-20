package sudoku_solver

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type Puzzle [9][9]int

type PuzzleExample struct {
	Puzzle     string `json:"puzzle"`
	Solution   string `json:"solution"`
	Difficulty string `json:"difficulty"`
}

// pass -1 for all, or pass n for at most n random puzzles
func Getdata(amount int) []PuzzleExample {
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
		rand.Shuffle(len(puzzle_examples), func(i, j int) {
			puzzle_examples[i], puzzle_examples[j] = puzzle_examples[j], puzzle_examples[i]
		})
		puzzle_examples = puzzle_examples[0:amount]
	}

	return puzzle_examples
}

func RawToPuzzle(p string) Puzzle {
	if len(p) < 81 {
		log.Fatal("Only 81 char strings can be processed!")
	}

	var po Puzzle
	var i int

	for j := 0; j < len(po); j++ {
		for k := 0; k < len(po[j]); k++ {
			digitChar := p[i]

			if digitChar < '0' || digitChar > '9' {
				log.Fatal("raw puzzle must contain only integers!")
			}

			po[j][k] = int(digitChar)

			i++
		}
	}

	return po
}

func PuzzleToRaw(p Puzzle) string {
	var po string

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(p[i]); j++ {
			po += string(p[i][j])
		}
	}

	return po
}
