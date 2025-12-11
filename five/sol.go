package five

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func checkRanges(idRanges [][]int, newRange []int) ([][]int) {
	// Try to consolidate
	for i, curRange := range idRanges {
		left := curRange[0]
		right := curRange[1]

		lIn := newRange[0] >= left && newRange[0] <= right
		rIn := newRange[1] >= left && newRange[1] <= right

		if lIn && rIn {
			return idRanges
		}

		if newRange[0] <= left && newRange[1] >= right {
			idRanges[i] = []int{newRange[0], newRange[1]}
			return idRanges
		}

		if lIn && newRange[1] >= right {
			idRanges[i] = []int{left, newRange[1]}
			return idRanges
		}

		if newRange[0] <= left && rIn {
			idRanges[i] = []int{newRange[0], right}
			return idRanges
		}
	}
	// else, append to array
	return append(idRanges, newRange)
}

func One() {
	file, err := os.Open("./five/day5-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}

	scanner := bufio.NewScanner(file)

	var idRanges [][]int
	var ingredients []int

	brk := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			brk = true
			continue
		}

		if !brk {
			ids := strings.Split(line, "-")
			id1, _ := strconv.Atoi(ids[0])
			id2, _ := strconv.Atoi(ids[1])
			idNums := []int{id1, id2}

			if len(idRanges) == 0 {
				idRanges = append(idRanges, idNums)
			} else {
				idRanges = checkRanges(idRanges, idNums)
			}
		} else {
			ingredientId, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredientId)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file line: %s\n", err.Error())
	}

	numFresh := 0
	for _, ingredientId := range ingredients {
		for _, idRange := range idRanges {
			if ingredientId >= idRange[0] && ingredientId <= idRange[1] {
				numFresh++
				break
			}
		}
	}

	log.Printf("Number of fresh ingredients: %d\n", numFresh)
}
