package day_03

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"regexp"
	"strconv"
	"strings"
)

// eval receives a row as input and evaluates the valid multiplications
func eval(row string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0
	matches := re.FindAllString(row, -1)

	for _, match := range matches {
		sum += mul(match)
	}

	return sum
}

// evalAdvanced receives a row as input and evaluates the valid multiplications with additional switch logic
func evalAdvanced(row string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true
	matches := re.FindAllString(row, -1)

	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			sum += mul(match)
		}
	}

	return sum
}

// mul receives a mul instruction and returns the multiplication result
func mul(str string) int {
	values := strings.Split(str[4:len(str)-1], ",")
	return utils.Atoi(values[0]) * utils.Atoi(values[1])
}

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
	sum := 0

	for _, row := range input {
		sum += eval(row)
	}

	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	joined := ""

	for _, row := range input {
		joined += row
	}

	return strconv.Itoa(evalAdvanced(joined))
}
