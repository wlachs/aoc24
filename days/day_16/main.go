package day_16

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"math"
	"slices"
	"strconv"
)

// vecDir holds a position vector and a facing
type vecDir struct {
	vec, dir types.Vec2
}

// memory struct for optimizing the recursive algorithm
type memory struct {
	visited []types.Vec2
}

// start gets the starting point of the map
func start(m map[types.Vec2]int32) types.Vec2 {
	for vec2, i := range m {
		if i == 'S' {
			return vec2
		}
	}
	panic("start not found")
}

// end gets the ending point of the map
func end(m map[types.Vec2]int32) types.Vec2 {
	for vec2, i := range m {
		if i == 'E' {
			return vec2
		}
	}
	panic("end not found")
}

// costs calculates the cost to travel from the starting node to every reachable node of the map
func costs(m map[types.Vec2]int32) map[vecDir]int {
	s := vecDir{start(m), types.Vec2{X: 1}}
	r := map[vecDir]int{}
	r[s] = 0
	visit := []vecDir{s}
	add := func(cur, next vecDir, cost int) {
		if n, ok := r[next]; m[next.vec] != '#' && (!ok || r[cur]+cost < n) {
			r[next] = r[cur] + cost
			visit = append(visit, next)
		}
	}
	for len(visit) > 0 {
		cur := visit[0]
		visit = visit[1:]
		add(cur, vecDir{cur.vec.Add(&cur.dir), cur.dir}, 1)
		add(cur, vecDir{cur.vec, cur.dir.RotateLeft()}, 1000)
		add(cur, vecDir{cur.vec, cur.dir.RotateRight()}, 1000)
	}
	return r
}

// minCost calculates the minimum cost to reach the given vector
func minCost(c map[vecDir]int, vec types.Vec2) int {
	mc := math.MaxInt
	for key, cost := range c {
		if key.vec == vec {
			mc = min(mc, cost)
		}
	}
	return mc
}

// hasCost checks if the cost matrix takes a specific value at a given vector
func hasCost(c map[vecDir]int, vec types.Vec2, val int) bool {
	for key, cost := range c {
		if key.vec == vec && cost == val {
			return true
		}
	}
	return false
}

// paths calculates the points on the optimal path(s) from the given vector based on the cost map
func paths(m map[types.Vec2]int32, c map[vecDir]int, e vecDir, cost int, nodes *memory) {
	appendList := func(l types.Vec2) {
		if !slices.Contains(nodes.visited, l) {
			nodes.visited = append(nodes.visited, l)
		}
	}
	leftDir := e.dir.RotateLeft()
	rightDir := e.dir.RotateRight()
	straight := e.vec.Subtract(&e.dir)
	left := e.vec.Subtract(&leftDir)
	right := e.vec.Subtract(&rightDir)
	if hasCost(c, straight, cost-1) && !slices.Contains(nodes.visited, straight) {
		paths(m, c, vecDir{straight, e.dir}, cost-1, nodes)
	}
	if hasCost(c, left, cost-1001) && !slices.Contains(nodes.visited, left) {
		paths(m, c, vecDir{left, leftDir}, cost-1001, nodes)
	}
	if hasCost(c, right, cost-1001) && !slices.Contains(nodes.visited, right) {
		paths(m, c, vecDir{right, rightDir}, cost-1001, nodes)
	}
	appendList(e.vec)
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
	m := utils.ParseInputToMap(input)
	return strconv.Itoa(minCost(costs(m), end(m)))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	c := costs(m)
	e := end(m)
	memo := memory{}
	paths(m, c, vecDir{e, types.Vec2{X: 1}}, minCost(c, e), &memo)
	paths(m, c, vecDir{e, types.Vec2{Y: -1}}, minCost(c, e), &memo)
	return strconv.Itoa(len(memo.visited))
}
