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

// regionData contains the area and perimeter / sides of a field
type regionData struct {
	area             uint64
	perimeterOrSides uint64
}

// fieldAndOrientation contains a pair of vectors indicating position and facing
type fieldAndOrientation struct {
	field       types.Vec2
	orientation types.Vec2
}

// regionCtx holds the calculation context
type regionCtx struct {
	fields                []types.Vec2
	fieldsWithOrientation []fieldAndOrientation
	perimeterOrSides      uint64
}

// calculate calculates the area and perimeter of each region
func calculate(m map[types.Vec2]regionKey, sides bool) []regionData {
	var res []regionData
	processed := make([]types.Vec2, 0, len(m))
	for pos, key := range m {
		if slices.Contains(processed, pos) {
			continue
		}
		var ctx regionCtx
		find(m, pos, pos, key, &ctx, sides)
		processed = append(processed, ctx.fields...)
		res = append(res, regionData{uint64(len(ctx.fields)), ctx.perimeterOrSides})
	}
	return res
}

// find executes a recursive search starting from pos in m and fills the fields with similarly marked neighbouring coordinates.
// If the sides property is set then instead of the perimeter, the number of sides is calculated
func find(m map[types.Vec2]regionKey, prev, pos types.Vec2, key regionKey, ctx *regionCtx, sides bool) {
	if m[pos] != key {
		if !sides {
			ctx.perimeterOrSides++
			return
		}
		dir := pos.Subtract(&prev)
		if !slices.Contains(ctx.fieldsWithOrientation, fieldAndOrientation{prev, dir}) {
			ctx.perimeterOrSides++
			leftDir := dir.RotateLeft()
			rightDir := dir.RotateRight()
			for pIn, pOut := prev, pos; m[pIn] == key && m[pOut] != key; pIn, pOut = pIn.Add(&leftDir), pOut.Add(&leftDir) {
				ctx.fieldsWithOrientation = append(ctx.fieldsWithOrientation, fieldAndOrientation{pIn, dir})
			}
			for pIn, pOut := prev.Add(&rightDir), pos.Add(&rightDir); m[pIn] == key && m[pOut] != key; pIn, pOut = pIn.Add(&rightDir), pOut.Add(&rightDir) {
				ctx.fieldsWithOrientation = append(ctx.fieldsWithOrientation, fieldAndOrientation{pIn, dir})
			}
		}
		return
	} else if slices.Contains(ctx.fields, pos) {
		return
	}
	ctx.fields = append(ctx.fields, pos)
	for _, nextPos := range pos.Around() {
		find(m, pos, nextPos, key, ctx, sides)
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
	for _, r := range calculate(m, false) {
		sum += r.area * r.perimeterOrSides
	}
	return strconv.FormatUint(sum, 10)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	sum := uint64(0)
	for _, r := range calculate(m, true) {
		sum += r.area * r.perimeterOrSides
	}
	return strconv.FormatUint(sum, 10)
}
