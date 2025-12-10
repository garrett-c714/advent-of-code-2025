package two

import (
	"os"
	"log"
	"bufio"
	"strings"
	"math"
	"strconv"
)

func One() {
	file, err := os.Open("./two/day2-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	
	rangesLine := scanner.Text()

	rangesList := strings.Split(rangesLine, ",")

	var ranges [][]string
	for _, entry := range rangesList {
		ends := strings.Split(entry, "-")

		ranges = append(ranges, ends)
	}

	for i := range ranges {
		log.Printf("[%s, %s]\n", ranges[i][0], ranges[i][1])
	}

	answer := 0
	for _, entry:= range ranges {
		if len(entry[0]) == len(entry[1]) && len(entry[0]) % 2 == 1 {
			continue
		}
		leftBound, _ := strconv.Atoi(entry[0])
		rightBound, _ := strconv.Atoi(entry[1])

		leftLen := len(entry[0])
		rightLen := len(entry[1])
		for leftLen <= rightLen {
			if leftLen % 2 == 1 {
				leftLen += 1
				continue
			}
			half := leftLen / 2
			leftInt := math.Pow(10, float64(half - 1))
			rightInt := math.Pow(10, float64(half))

			for i := leftInt; i < rightInt; i++ {
				halfStr := strconv.Itoa(int(i))
				check, _ := strconv.Atoi(halfStr + halfStr)
				if check >= leftBound && check <= rightBound {
					answer += check
				}
			}


			leftLen += 1
		}
	}

	log.Printf("Answer: %d", answer)

}

func Two() {
	file, err := os.Open("./two/day2-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	
	rangesLine := scanner.Text()

	rangesList := strings.Split(rangesLine, ",")

	var ranges [][]string
	for _, entry := range rangesList {
		ends := strings.Split(entry, "-")

		ranges = append(ranges, ends)
	}

	for i := range ranges {
		log.Printf("[%s, %s]\n", ranges[i][0], ranges[i][1])
	}

	answer := 0
	for _, entry:= range ranges {
		if len(entry[0]) == len(entry[1]) && len(entry[0]) % 2 == 1 {
			continue
		}
		leftBound, _ := strconv.Atoi(entry[0])
		rightBound, _ := strconv.Atoi(entry[1])

		leftLen := len(entry[0])
		rightLen := len(entry[1])
		for leftLen <= rightLen {
			if leftLen % 2 == 1 {
				for i:=1; i<10; i++ {
					digit := strconv.Itoa(i)
					numStr := strings.Repeat(digit, leftLen)
					check, _ := strconv.Atoi(numStr)
					if check >= leftBound && check <= rightBound {
						answer += check
					}
				}

				leftLen += 1
				continue
			}
			half := leftLen / 2
			leftInt := math.Pow(10, float64(half - 1))
			rightInt := math.Pow(10, float64(half))

			for i := leftInt; i < rightInt; i++ {
				halfStr := strconv.Itoa(int(i))
				check, _ := strconv.Atoi(halfStr + halfStr)
				if check >= leftBound && check <= rightBound {
					answer += check
				}
			}


			leftLen += 1
		}
	}

	log.Printf("Answer: %d", answer)

}
