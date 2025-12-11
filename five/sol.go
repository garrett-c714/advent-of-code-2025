package five

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"slices"
)

func consolidateRanges(idRanges [][]int, newRange []int) ([][]int) {
	// Try to consolidate
	leftIn := -1
	rightIn := -1
	for i, curRange := range idRanges {
		// Swallow
		if newRange[0] <= curRange[0] && newRange[1] >= curRange[1] {
			idRanges[i][0] = -1
			idRanges[i][1] = -1
			continue
		}

		if newRange[0] >= curRange[0] && newRange[0] <= curRange[1] {
			leftIn = i
		}

		if newRange[1] >= curRange[0] && newRange[1] <= curRange[1] {
			rightIn = i
		}

		// Is Swallowed
		if leftIn >= 0 && rightIn >= 0 && leftIn == rightIn {
			return idRanges
		}
	}

	if rightIn >= 0 && leftIn >= 0 {
		newLeft := idRanges[leftIn][0]
		newRight := idRanges[rightIn][1]

		higherIndex := max(leftIn, rightIn)
		lowerIndex := min(leftIn, rightIn)

		idRanges = slices.Delete(idRanges, higherIndex, higherIndex+1)
		idRanges = slices.Delete(idRanges, lowerIndex, lowerIndex+1)

		idRanges = append(idRanges, []int{newLeft, newRight})
	} else if rightIn >= 0 {
		idRanges[rightIn] = []int{newRange[0], idRanges[rightIn][1]}
	} else if leftIn >= 0 {
		idRanges[leftIn] = []int{idRanges[leftIn][0], newRange[1]}
	} else {
		idRanges = append(idRanges, newRange)
	}

	i := 0
	for i < len(idRanges) {
		if idRanges[i][0] == -1 && idRanges[i][1] == -1 {
			idRanges = slices.Delete(idRanges, i, i+1)
			continue
		}

		i++
	}

	return idRanges
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
				idRanges = consolidateRanges(idRanges, idNums)
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
			}
		}
	}

	log.Printf("Number of fresh ingredients: %d\n", numFresh)
}

func Two() {
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
				idRanges = consolidateRanges(idRanges, idNums)
			}
		} else {
			ingredientId, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredientId)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file line: %s\n", err.Error())
	}

	totalFreshIds := 0
	for _, curRange := range idRanges {
		totalFreshIds += (curRange[1] - curRange[0]) + 1
	}

	log.Printf("Total fresh IDs: %d\n", totalFreshIds)
}

