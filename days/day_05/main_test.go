package day_05_test

import (
	"github.com/wlchs/aoc24/days/day_05"
	"github.com/wlchs/aoc24/internal"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_1_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_1.txt")
	result := day_05.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_2_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_2.txt")
	result := day_05.Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}
