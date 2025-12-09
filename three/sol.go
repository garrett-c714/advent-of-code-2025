package three

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"sort"
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

		tensIdx, tens := getMax(batteries[:rowLen - 1])
		_, ones := getMax(batteries[tensIdx + 1:])

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
	for i:=len(nums) - 1; i >= 0; i-- {
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

		var enabledBats []Occurrence
		target := 9
		for len(enabledBats) < 12 {
			log.Printf("Getting %d\n", target)
			occurences := getOccurences(batteries, target, 
				12 - len(enabledBats))
			enabledBats = append(enabledBats, occurences...)
			target -= 1
		}

		sort.Slice(enabledBats, func (i, j int) bool {
			return enabledBats[i].Index < enabledBats[j].Index
		})

		var enabledString = ""
		for _, occ := range enabledBats {
			strJolt  := strconv.Itoa(occ.Value)
			enabledString += strJolt
		}

		rowJoltage, _ := strconv.Atoi(enabledString)
		log.Printf("Calculated Joltage: %s\n", enabledString)
		//log.Printf("Length of calc-ed joltage: %d", len(enabledString))
		joltage += rowJoltage
	}

	log.Printf("Max Joltage: %d", joltage)

}

