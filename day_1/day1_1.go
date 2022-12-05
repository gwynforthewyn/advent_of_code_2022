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
	currentMaxCalories := currentElfSubtotal

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			currentElfSubtotal = 0
		} else {
			calories, err := strconv.ParseInt(line, 10, 32)

			if err != nil {
				panic(err)
			}

			currentElfSubtotal += calories

			if currentElfSubtotal > currentMaxCalories {
				currentMaxCalories = currentElfSubtotal
			}
		}
	}

	fmt.Println(currentMaxCalories)

}
