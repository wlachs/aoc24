package day_16_test

import (
	"github.com/wlchs/aoc24/days/day_16"
	"github.com/wlchs/aoc24/internal"
	"testing"
)

func TestPartOneA(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_1a_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_1a.txt")
	result := day_16.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartOneB(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_1b_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_1b.txt")
	result := day_16.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwoA(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_2a_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_2a.txt")
	result := day_16.Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwoB(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_2b_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_2b.txt")
	result := day_16.Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}
