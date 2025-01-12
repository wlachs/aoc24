package day_20

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
)

// key is a vector pair to be used for the Floyd-Warshall algorithm
type key struct {
	v1, v2 types.Vec2
}

// s finds the start position in the map
func s(m map[types.Vec2]int32) types.Vec2 {
	vec, _ := utils.FindUniqueInMap(m, 'S')
	return vec
}

// e finds the end position in the map
func e(m map[types.Vec2]int32) types.Vec2 {
	vec, _ := utils.FindUniqueInMap(m, 'E')
	return vec
}

// dijkstra finds the shortest path between the two positions if possible
func dijkstra(m map[types.Vec2]int32, start, end types.Vec2) []types.Vec2 {
	costMap := map[types.Vec2]int{}
	for vec2 := range m {
		costMap[vec2] = math.MaxInt
	}
	costMap[start] = 0
	visit := []types.Vec2{start}
	for len(visit) > 0 {
		h := visit[0]
		visit = visit[1:]
		for _, next := range h.Around() {
			if c, ok := m[next]; ok && c != '#' && costMap[h]+1 < costMap[next] {
				visit = append(visit, next)
				costMap[next] = costMap[h] + 1
				if next == end {
					return path(costMap, start, end)
				}
			}
		}
	}
	panic("no path found")
}

// wallDijkstra finds the shortest paths starting from a given source
func wallDijkstra(m map[types.Vec2]int32, start types.Vec2) map[types.Vec2]int {
	costMap := map[types.Vec2]int{}
	for vec2 := range m {
		costMap[vec2] = math.MaxInt
	}
	costMap[start] = 0
	visit := []types.Vec2{start}
	for len(visit) > 0 {
		h := visit[0]
		visit = visit[1:]
		for _, next := range h.Around() {
			if _, ok := m[next]; ok && costMap[h]+1 < costMap[next] {
				costMap[next] = costMap[h] + 1
				visit = append(visit, next)
			}
		}
	}
	return costMap
}

// allDijkstra runs the Dijkstra algorithm on each node of the graph
func allDijkstra(m map[types.Vec2]int32, size int) [][]int {
	var pathFields []int
	for vec2, c := range m {
		if c == '.' || c == 'S' || c == 'E' {
			pathFields = append(pathFields, vecToNode(vec2, size))
		}
	}
	nodes := len(m)
	dist := make([][]int, nodes)
	for i := range nodes {
		dist[i] = make([]int, nodes)
		for j := range nodes {
			if i == j {
				dist[i][j] = 0
			} else if areNeighbors(i, j, size) && notBothOnPath(pathFields, i, j) {
				dist[i][j] = 1
			} else {
				dist[i][j] = 21
			}
		}
	}
	for source := range nodes {
		q := make([]int, nodes)
		for v := range nodes {
			q = append(q, v)
		}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, v := range neighbors(u, size) {
				if dist[u][v] == 1 && dist[source][u]+1 < dist[source][v] {
					dist[source][v] = dist[source][u] + 1
				}
			}
		}
	}
	return dist
}

// vecToNode converts a vector to a graph node
func vecToNode(vec types.Vec2, size int) int {
	return vec.X*size + vec.Y
}

// areNeighbors checks whether two nodes are neighbors in the graph
func areNeighbors(i, j, size int) bool {
	for _, n := range neighbors(i, size) {
		if n == j {
			return true
		}
	}
	return false
}

// neighbors returns the direct neighbors of a graph node
func neighbors(n, size int) []int {
	res := make([]int, 0, 4)
	if n%size != 0 {
		res = append(res, n-1)
	}
	if n%size != size-1 {
		res = append(res, n+1)
	}
	if n > size {
		res = append(res, n-size)
	}
	if n < (size-1)*size {
		res = append(res, n+size)
	}
	return res
}

// notBothOnPath makes sure that the two given characters are not both path characters i.e. either '.', 'S' or 'E'
func notBothOnPath(pathFields []int, f1, f2 int) bool {
	c := 0
	for _, field := range pathFields {
		if field == f1 || field == f2 {
			c++
			if c == 2 {
				return false
			}
		}
	}
	return true
}

// path finds one of the shortest paths on the map from the start to the end position
func path(costMap map[types.Vec2]int, start types.Vec2, end types.Vec2) []types.Vec2 {
	cur := end
	p := []types.Vec2{cur}
	for cur != start {
		for _, vec2 := range cur.Around() {
			if costMap[vec2] < costMap[cur] {
				cur = vec2
			}
		}
		p = append([]types.Vec2{cur}, p...)
	}
	return p
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
	p := dijkstra(m, s(m), e(m))
	count := 0
	for i, start := range p[:len(p)-5] {
		for j, end := range p[i+4:] {
			d := start.Subtract(&end)
			if d.X*d.Y == 0 && utils.Abs(d.X+d.Y) == 2 && j+2 >= 20 {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	p := dijkstra(m, s(m), e(m))
	count := 0
	size := len(p)
	for i, start := range p {
		d := wallDijkstra(m, start)
		for j, end := range p[i:] {
			wallPath := d[end]
			if wallPath <= 20 && j-wallPath >= 100 {
				count++
			}
		}
		fmt.Println(i, float64(i)/float64(size)*100, count)
	}
	return strconv.Itoa(count)
}
