package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"math-skills/mathskills"
)

func main() {
	numbers := []float64{}
	scanner := bufio.NewScanner(os.Stdin)
	windowSize := 2 // Adjust window size as needed

	for scanner.Scan() {
		lines := scanner.Text()
		numbs, err := strconv.ParseFloat(lines, 64)
		if err != nil {
			log.Fatalf("Error during conversion")
		}

		numbers = append(numbers, numbs)
		if len(numbers) > windowSize {
			numbers = numbers[1:]
		}

		if len(numbers) == windowSize {
			lower, upper := mathskills.Nextrange(numbers)
			fmt.Printf("%d - %d\n", lower, upper)
		}
	}
}
