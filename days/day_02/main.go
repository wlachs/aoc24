package day_02

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := 0

	for _, report := range input {
		valid := true
		values := utils.ToIntSlice(strings.Split(report, " "))
		sign := math.Signbit(float64(values[1] - values[0]))
		for i := 1; i < len(values); i++ {
			diff := values[i] - values[i-1]
			absDiff := utils.Abs(diff)

			if math.Signbit(float64(diff)) != sign || absDiff < 1 || absDiff > 3 {
				valid = false
			}
		}

		if valid {
			count++
		}
	}

	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
