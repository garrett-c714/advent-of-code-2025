package three

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
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

