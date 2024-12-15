package day_14

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"regexp"
	"slices"
	"strconv"
)

// robot struct containing position and velocity
type robot struct {
	p, v types.Vec2
}

// quadrant returns the ID of the quadrant where the robot current is
// If the robot is exactly on the border between two quadrants, return false
func (r *robot) quadrant(width, height int) (int, bool) {
	w, h := width/2, height/2
	if r.p.X == w || r.p.Y == h {
		return -1, false
	}
	return 2*(r.p.X/(w+1)) + r.p.Y/(h+1), true
}

// move simulates a robot movement
func (r *robot) move(width, height int) {
	r.p.X = utils.Mod(r.p.X+r.v.X, width)
	r.p.Y = utils.Mod(r.p.Y+r.v.Y, height)
}

// parseRobots parses the input and returns a slice of robots
func parseRobots(input []string) []robot {
	robots := make([]robot, 0, len(input))
	re := regexp.MustCompile(`.*=(?P<pX>\d+),(?P<pY>\d+).*=(?P<vX>.*),(?P<vY>.*)`)
	for _, s := range input {
		match := re.FindStringSubmatch(s)
		p := types.Vec2{X: utils.Atoi(match[1]), Y: utils.Atoi(match[2])}
		v := types.Vec2{X: utils.Atoi(match[3]), Y: utils.Atoi(match[4])}
		robots = append(robots, robot{p, v})
	}
	return robots
}

// hasNLargeConnectedSubGraph searches the robot graph for an n large connected subgraph
func hasNLargeConnectedSubGraph(robots []robot, n int) bool {
	for _, r := range robots {
		visitedElements := make([]types.Vec2, 0, n)
		elementsToVisit := []types.Vec2{r.p}
		for len(elementsToVisit) > 0 {
			for _, vec := range elementsToVisit[0].Around() {
				if checkRobotAt(robots, vec) && !slices.Contains(visitedElements, vec) {
					elementsToVisit = append(elementsToVisit, vec)
				}
			}
			visitedElements = append(visitedElements, elementsToVisit[0])
			elementsToVisit = elementsToVisit[1:]
			if len(visitedElements) >= n {
				return true
			}
		}
	}
	return false
}

// checkRobotAt checks if there is a robot at the given vector
func checkRobotAt(robots []robot, vec types.Vec2) bool {
	for _, r := range robots {
		if r.p == vec {
			return true
		}
	}
	return false
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input, 101, 103))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input, 101, 103))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string, width, height int) string {
	robots := parseRobots(input)
	quadrants := make([]uint64, 4)
	for range 100 {
		for i := range robots {
			robots[i].move(width, height)
		}
	}
	for _, r := range robots {
		if q, ok := r.quadrant(width, height); ok {
			quadrants[q]++
		}
	}
	return strconv.FormatUint(quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3], 10)
}

// Part2 solves the second part of the exercise
func Part2(input []string, width, height int) string {
	robots := parseRobots(input)
	round := 0
	for {
		if hasNLargeConnectedSubGraph(robots, 30) {
			return strconv.Itoa(round)
		}
		for i := range robots {
			robots[i].move(width, height)
		}
		round++
	}
}
