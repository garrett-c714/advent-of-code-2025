package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func one() {
	file, err := os.Open("./day1-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()

		//log.Printf("Starting at: %d", dial)

		dir := string(line[0])
		shift, _ := strconv.Atoi(line[1:])

		if dir == "R" {
			dial += shift % 100
		} else {
			dial -= shift % 100
		}

		if dial > 99 {
			dial = dial % 100
		} else if dial < 0 {
			dial = 100 + dial
		}

		//log.Printf("Rotating %s %d --> %d", dir, shift, dial)

		if dial == 0 {
			password++
		}

		//log.Println()
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file line: %s", err.Error())
	}

	log.Printf("Password: %d", password)
}

func two() {
	file, err := os.Open("./day1-input.txt")
	if err != nil {
		log.Fatal("Could not open input file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()

		//log.Printf("Starting at: %d", dial)

		dir := string(line[0])
		shift, _ := strconv.Atoi(line[1:])

		for shift > 0 {
			if dir == "L" {
				dial -= 1
			} else {
				dial += 1
			}

			if dial == -1 {
				dial = 99
			}

			if dial == 100 {
				dial = 0
			}

			if dial == 0 {
				password++
			}
			
			shift -= 1
		}

	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file line: %s", err.Error())
	}

	log.Printf("Password: %d", password)
}


func main() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "one":
		one()
	case "two":
		two()
	}

}
