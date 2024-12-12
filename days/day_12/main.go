package day_12

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
)

// regionKey is an alias for int32, just to make the code more readable
type regionKey = int32

// regionData contains the area and perimeter of a fields
type regionData struct {
	area      uint64
	perimeter uint64
}

// regionCtx holds the calculation context
type regionCtx struct {
	fields    []types.Vec2
	perimeter uint64
}

// calculate calculates the area and perimeter of each regionData
func calculate(m map[types.Vec2]regionKey) []regionData {
	var res []regionData
	regions := map[regionKey][][]types.Vec2{}
	processed := make([]types.Vec2, 0, len(m))
	for pos, key := range m {
		if slices.Contains(processed, pos) {
			continue
		}
		var ctx regionCtx
		find(m, pos, key, &ctx)
		if _, ok := regions[key]; !ok {
			regions[key] = [][]types.Vec2{}
		}
		regions[key] = append(regions[key], ctx.fields)
		processed = append(processed, ctx.fields...)
		res = append(res, regionData{uint64(len(ctx.fields)), ctx.perimeter})
	}
	return res
}

// find executes a recursive search starting from pos in m and fills the fields with similarly marked neighbouring coordinates
func find(m map[types.Vec2]regionKey, pos types.Vec2, key regionKey, ctx *regionCtx) {
	if m[pos] != key {
		ctx.perimeter++
		return
	} else if slices.Contains(ctx.fields, pos) {
		return
	}
	ctx.fields = append(ctx.fields, pos)
	for _, nextPos := range pos.Around() {
		find(m, nextPos, key, ctx)
	}
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
	sum := uint64(0)
	for _, r := range calculate(m) {
		sum += r.area * r.perimeter
	}
	return strconv.FormatUint(sum, 10)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
