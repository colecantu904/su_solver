package sudoku_solver

import (
	"log"
	"encoding/json"
	"os"
	"math/rand"
)

type Puzzle struct {
	Puzzle string `json:"puzzle"`
	Solution string `json:"solution"`
	Difficulty string `json:"difficulty"`
}

// pass -1 for all, or pass n for at most n random puzzles
func Getdata( amount int ) []Puzzle {
	var puzzles []Puzzle

	data, err := os.ReadFile("testdata/sudoku.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &puzzles)
	if err != nil {
		log.Fatal(err)
	}

	if amount != -1 {
		rand.Shuffle(len(puzzles), func( i, j int ) {
			puzzles[i], puzzles[j] = puzzles[j], puzzles[i]
		})
		puzzles = puzzles[0:amount]
	}

	return puzzles
}
