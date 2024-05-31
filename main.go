package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"math-skills/mathskills"
)

// This Go program reads numeric values from a file, calculates average, median, variance, and standard deviation.
func main() {
	if len(os.Args) != 2 {
		return
	}
	args, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}
	data := string(args)

	if !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Use a txt file")
		os.Exit(1)
	}

	value := []float64{}
	g := strings.Split(data, "\n")
	for _, str := range g {

		str = strings.TrimSpace(str)

		if str == "" {
			continue
		}

		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Error converting the string to float", err)
			os.Exit(1)
		}

		value = append(value, num)
	}
	if len(value) == 0 {
		return
	}
	//sorts the data in the data.txt file
	sort.Float64s(value)
	//calling the functions to be used in calculating the content of the data.txt file
	average := mathskills.Average(value)
	median := mathskills.Median(value)
	variance := mathskills.Variance(value)
	deviation := mathskills.Deviation(value)

	fmt.Printf("Average: %.0f\n", math.Round(average))
	fmt.Printf("Median: %.0f\n", math.Round(median))
	fmt.Printf("Variance: %.0f\n", math.Round(variance))
	fmt.Printf("Standard Deviation: %.0f\n", math.Round(deviation))
}
