package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input/actual_1_inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	currentElfSubtotal := int64(0)

	var subTotals []int64

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			subTotals = append(subTotals, currentElfSubtotal)
			currentElfSubtotal = 0
		} else {
			calories, err := strconv.ParseInt(line, 10, 32)

			if err != nil {
				panic(err)
			}

			currentElfSubtotal += calories
		}
	}

	// This only works if the first element in the subTotals list isn't coincidentally one of the 3 largest.
	first := subTotals[0]
	second := subTotals[0]
	third := subTotals[0]

	for _, subTotal := range subTotals {
		if subTotal > first {
			first = subTotal
			continue
		}

		if subTotal > second && subTotal != first {
			second = subTotal
			continue
		}

		if subTotal > third && subTotal != first && subTotal != second {
			third = subTotal
			continue
		}
	}
	fmt.Println(subTotals)
	fmt.Println(third + second + first)

}
