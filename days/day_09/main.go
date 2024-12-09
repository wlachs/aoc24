package day_09

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"strconv"
)

// buildDiskMap parses the input and builds the disk map as int slice
func buildDiskMap(input string) []int {
	var disk []int
	sector := 0

	for i, char := range input {
		num := utils.Atoi(string(char))

		if i%2 == 0 {
			for range num {
				disk = append(disk, sector)
			}
			sector++

		} else {
			for range num {
				disk = append(disk, -1)
			}
		}
	}

	return disk
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
	disk := buildDiskMap(input[0])
	checksum := 0

	for x := 0; x < len(disk); x++ {
		if disk[x] == -1 {
			for y := len(disk) - 1; y > x; y-- {
				if disk[y] != -1 {
					disk[x], disk[y] = disk[y], disk[x]
					break
				}
			}
		}

		if disk[x] != -1 {
			checksum += x * disk[x]
		}
	}

	return strconv.Itoa(checksum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	disk := buildDiskMap(input[0])
	checksum := 0
	block := -1
	blockLength := 0

	for x := len(disk) - 1; x >= 0; x-- {
		if disk[x] == block {
			blockLength++
		} else {
			if block != -1 {
				freeSpace := 0
				for y := 0; y <= x; y++ {
					if disk[y] == -1 {
						freeSpace++
					} else {
						freeSpace = 0
					}
					if freeSpace == blockLength {
						for z := range freeSpace {
							disk[y-z], disk[x+z+1] = disk[x+z+1], disk[y-z]
						}
						break
					}
				}
			}

			block = disk[x]
			blockLength = 1
		}
	}

	for x, file := range disk {
		if file != -1 {
			checksum += x * file
		}
	}

	return strconv.Itoa(checksum)
}
