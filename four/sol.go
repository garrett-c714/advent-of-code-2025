package four

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func incNeighbors(adjacencies [][]int, i int, j int,
	maxRow int, maxCol int) {
	hasTop := i > 0
	hasLeft := j > 0
	hasBot := i < maxRow
	hasRight := j < maxCol

	if hasTop {
		adjacencies[i-1][j]++
		if hasLeft {
			adjacencies[i-1][j-1]++
		}
		if hasRight {
			adjacencies[i-1][j+1]++
		}
	}

	if hasLeft {
		adjacencies[i][j-1]++
	}

	if hasRight {
		adjacencies[i][j+1]++
	}

	if hasBot {
		adjacencies[i+1][j]++
		if hasLeft {
			adjacencies[i+1][j-1]++
		}
		if hasRight {
			adjacencies[i+1][j+1]++
		}
	}
}

func One() {
	file, err := os.Open("./four/day4-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}

	scanner := bufio.NewScanner(file)

	var adjacencies [][]int
	var rowCells [][]string

	for scanner.Scan() {
		line := scanner.Text()

		row := strings.Split(line, "")
		rowCells = append(rowCells, row)
		adjacencies = append(adjacencies, make([]int, len(row)))

	}
	if scanner.Err(); err != nil {
		log.Printf("Error scanning line in input: %s\n", err.Error())
	}

	numRows := len(rowCells)
	rowLen := len(rowCells[0])

	for i, _ := range rowCells {
		for j, cell := range rowCells[i] {
			if cell == "@" {
				incNeighbors(
					adjacencies,
					i,
					j,
					numRows-1,
					rowLen-1,
				)
			}
		}
	}

	numRolls := 0
	for i, _ := range rowCells {
		for j, cell := range rowCells[i] {
			if cell == "@" && adjacencies[i][j] < 4 {
				numRolls++
			}
		}
	}

	log.Printf("Number of Wrapping Paper Rolls: %d\n", numRolls)

}

func doRemoval(state [][]string) (numRemoved int) {
	adjacencies := make([][]int, len(state))
	for i := range len(state) {
		adjacencies[i] = make([]int, len(state[0]))
	}

	numRows := len(state)
	rowLen := len(state[0])

	for i, _ := range state {
		for j, cell := range state[i] {
			if cell == "@" {
				incNeighbors(
					adjacencies,
					i,
					j,
					numRows-1,
					rowLen-1,
				)
			}
		}
	}

	numRolls := 0
	for i, _ := range state {
		for j, cell := range state[i] {
			if cell == "@" && adjacencies[i][j] < 4 {
				state[i][j] = "."
				numRolls++
			}
		}
	}

	return numRolls
}

func Two() {
	file, err := os.Open("./four/day4-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}

	scanner := bufio.NewScanner(file)

	var rowCells [][]string

	for scanner.Scan() {
		line := scanner.Text()

		row := strings.Split(line, "")
		rowCells = append(rowCells, row)

	}
	if scanner.Err(); err != nil {
		log.Printf("Error scanning line in input: %s\n", err.Error())
	}

	removedRound := -1
	removedTotal := 0
	for removedRound != 0 {
		removedRound = doRemoval(rowCells)
		if removedRound > 0 {
			removedTotal += removedRound
		}
	}

	log.Printf("Total Number of Wrapping Paper Rolls: %d\n", removedTotal)
}

