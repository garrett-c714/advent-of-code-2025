package main 

import (
	"os"

	"advent-of-code-2025/one"
	"advent-of-code-2025/two"
	"advent-of-code-2025/three"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	switch os.Args[1] {
	case "one":
		switch os.Args[2] {
		case "one":
			one.One()
		case "two":
			one.Two()
		}
	case "two":
		switch os.Args[2] {
		case "one":
			two.One()
		}
	case "three":
		switch os.Args[2] {
		case "one":
			three.One()	
		case "two":
			three.Two()
		}
	}

}
