package day_11

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

// graph type holding the nodes and vertices
type graph = map[uint64][]uint64

// key contains a pair of node and depth to be used as key in the memory map
type key struct {
	node  uint64
	depth uint8
}

// memory type for accelerating the recursion
type memory = map[key]uint64

// buildGraph builds the stone graph based on the input
func buildGraph(nodes []uint64) graph {
	g := graph{}
	var processed []uint64
	for len(nodes) > 0 {
		var first uint64
		first, nodes = nodes[0], nodes[1:]
		processed = append(processed, first)
		nextNodes := findNextNodes(first)
		g[first] = nextNodes
		for _, node := range nextNodes {
			if !slices.Contains(processed, node) {
				nodes = append(nodes, node)
			}
		}
	}
	return g
}

// findNextNodes does the stone-splitting game logic to find the next nodes of the stone graph
func findNextNodes(stone uint64) []uint64 {
	if stone == 0 {
		return []uint64{1}
	} else if d := digits(stone); d%2 == 0 {
		exp := uint64(math.Pow10(int(d) / 2))
		return []uint64{stone / exp, stone % exp}
	} else {
		return []uint64{stone * 2024}
	}
}

// digits counts the digits of the given uint64
func digits(number uint64) uint8 {
	count := uint8(0)
	for number > 0 {
		number /= 10
		count++
	}
	return count
}

// search executes a DFS on the given graph starting from the given node and counts how many nodes are altogether visited
func search(node uint64, g graph, depth uint8, m memory) uint64 {
	if depth == 0 {
		return 1
	} else if count, ok := m[key{node, depth}]; ok {
		return count
	}
	count := uint64(0)
	for _, n := range g[node] {
		count += search(n, g, depth-1, m)
	}
	m[key{node, depth}] = count
	return count
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
	stones := utils.ToUInt64Slice(strings.Split(input[0], " "))
	g := buildGraph(stones)
	count := uint64(0)
	for _, stone := range stones {
		count += search(stone, g, 25, memory{})
	}
	return strconv.FormatUint(count, 10)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	stones := utils.ToUInt64Slice(strings.Split(input[0], " "))
	g := buildGraph(stones)
	count := uint64(0)
	for _, stone := range stones {
		count += search(stone, g, 75, memory{})
	}
	return strconv.FormatUint(count, 10)
}
