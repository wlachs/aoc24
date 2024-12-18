package day_17

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

// context holds the execution environment
type context struct {
	ic         uint8
	program    []uint8
	reg        [3]int
	initReg    [3]int
	output     []uint8
	operations [8]func(arg uint8)
}

// init initializes the application context
func (ctx *context) init(input []string) {
	for i := range 3 {
		ctx.reg[i] = utils.Atoi(strings.Split(input[i], ": ")[1])
		ctx.initReg[i] = ctx.reg[i]
	}
	program := strings.Split(strings.Split(input[4], ": ")[1], ",")
	ctx.program = utils.ToUInt8Slice(program)

	// adv: division -> A
	ctx.operations[0] = func(arg uint8) {
		ctx.reg[0] = ctx.reg[0] >> ctx.combo(arg)
		ctx.increment()
	}

	// bxl: bitwise xor (B ^ op)
	ctx.operations[1] = func(arg uint8) {
		ctx.reg[1] = ctx.reg[1] ^ int(ctx.literal(arg))
		ctx.increment()
	}

	// bst: combo % 8
	ctx.operations[2] = func(arg uint8) {
		ctx.reg[1] = ctx.combo(arg) & 7
		ctx.increment()
	}

	// jnz: jump
	ctx.operations[3] = func(arg uint8) {
		if ctx.reg[0] != 0 {
			ctx.ic = ctx.literal(arg)
		} else {
			ctx.increment()
		}
	}

	// bxc: bitwise xor (B ^ C)
	ctx.operations[4] = func(_ uint8) {
		ctx.reg[1] = ctx.reg[1] ^ ctx.reg[2]
		ctx.increment()
	}

	// out: write combo op
	ctx.operations[5] = func(arg uint8) {
		ctx.output = append(ctx.output, uint8(ctx.combo(arg)&7))
		ctx.increment()
	}

	// bdv: division -> B
	ctx.operations[6] = func(arg uint8) {
		ctx.reg[1] = ctx.reg[0] >> ctx.combo(arg)
		ctx.increment()
	}

	// cdv: division -> C
	ctx.operations[7] = func(arg uint8) {
		ctx.reg[2] = ctx.reg[0] >> ctx.combo(arg)
		ctx.increment()
	}
}

// literal gets the literal value of the operand
func (ctx *context) literal(op uint8) uint8 {
	return op
}

// combo gets the value of the given combo operand
func (ctx *context) combo(op uint8) int {
	if op <= 3 {
		return int(ctx.literal(op))
	} else if op == 7 {
		panic("invalid combo operand")
	}
	return ctx.reg[op-4]
}

// increment increments the instruction pointer by 2
func (ctx *context) increment() {
	ctx.ic += 2
}

// run executes the program with the given input
func (ctx *context) run() {
	ctx.reg[1], ctx.reg[2] = ctx.initReg[1], ctx.initReg[2]
	ctx.ic = 0
	ctx.output = nil
	for int(ctx.ic) < len(ctx.program)-1 {
		ctx.operations[ctx.program[ctx.ic]](ctx.program[ctx.ic+1])
	}
}

// formattedOutput formats the output slice of the program as comma-separated string
func (ctx *context) formattedOutput() string {
	o := make([]int, len(ctx.output))
	for i, u := range ctx.output {
		o[i] = int(u)
	}
	return strings.Join(utils.ToStringSlice(o), ",")
}

// rec recursively finds the matching A register value
func rec(bits, val int, ctx *context, n int) int {
	if n == -1 {
		return val
	}
	b3 := 3 * n
	mask := (7 << b3) ^ (1<<bits - 1)
	var values []int
	for i := range 8 {
		v := val&mask ^ (i << b3)
		ctx.reg[0] = v
		ctx.run()
		if len(ctx.output) > n && ctx.output[n] == (ctx.program[n]) {
			values = append(values, rec(bits, v, ctx, n-1))
		}
	}
	if len(values) == 0 {
		return math.MaxInt
	}
	return slices.Min(values)
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
	ctx := context{}
	ctx.init(input)
	ctx.run()
	return ctx.formattedOutput()
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	ctx := context{}
	ctx.init(input)
	l := len(ctx.program)
	r := rec(3*l, 1<<(3*l), &ctx, l-1)
	return strconv.Itoa(r)
}
