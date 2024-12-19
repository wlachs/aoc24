package day_19

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// f counts all possible ways to combine words to reach the desired outcome
func f(words []string, s string, m map[string]int) bool {
	if _, ok := m[s]; ok {
		return true
	}
	for i := len(s); i > 0; i-- {
		if slices.Contains(words, s[:i]) && f(words, s[i:], m) {
			m[s] += m[s[i:]]
		}
	}
	_, ok := m[s]
	return ok
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
	words := strings.Split(input[0], ", ")
	count := 0
	m := map[string]int{"": 1}
	for _, s := range input[2:] {
		if f(words, s, m) {
			count++
		}
	}
	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	words := strings.Split(input[0], ", ")
	count := 0
	m := map[string]int{"": 1}
	for _, s := range input[2:] {
		f(words, s, m)
		count += m[s]
	}
	return strconv.Itoa(count)
}
