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

