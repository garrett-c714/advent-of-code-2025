package three

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	//"slices"
)

func getMax(nums []int) (index int, value int) {
	idx := 0
	mx := nums[0]
	for i, val := range nums {
		if val > mx {
			idx = i
			mx = val
		}
	}
	return idx, mx
}

func One() {
	file, err := os.Open("./three/day3-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltage := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		rowLen := len(split)

		var batteries []int
		for _, s := range split {
			intVal, _ := strconv.Atoi(s)
			batteries = append(batteries, intVal)
		}

		tensIdx, tens := getMax(batteries[:rowLen-1])
		_, ones := getMax(batteries[tensIdx+1:])

		joltage += (10 * tens) + ones

	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file line: %s", err.Error())
	}

	log.Printf("Max Joltage: %d", joltage)

}

type Occurrence struct {
	Index int
	Value int
}

func getOccurences(nums []int, tgt int, max int) []Occurrence {
	var ocs []Occurrence
	for i := len(nums) - 1; i >= 0; i-- {
		if len(ocs) == max {
			break
		}
		if nums[i] == tgt {
			ocs = append(ocs, Occurrence{Index: i, Value: nums[i]})
		}
	}
	return ocs
}

func Two() {
	file, err := os.Open("./three/day3-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	joltage := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")

		log.Printf("Battery Bank: %s\n", line)

		var batteries []int
		for _, s := range split {
			intVal, _ := strconv.Atoi(s)
			batteries = append(batteries, intVal)
		}

		activeBats := ""
		fwdIdx := -1
		for range 12 {
			idx, digit := getMax(
				batteries[fwdIdx+1 : len(batteries) - (11 - len(activeBats))],
			)
			fwdIdx = fwdIdx + idx + 1
			activeBats += strconv.Itoa(digit)
		}

		log.Printf("Bank Joltage: %s\n", activeBats)
		maxJoltage, _ := strconv.Atoi(activeBats)
		joltage += maxJoltage

	}

	log.Printf("Max Joltage: %d", joltage)

}
