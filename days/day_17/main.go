package day_17

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"strings"
)

// context holds the execution environment
type context struct {
	ic         uint8
	program    []uint8
	reg        [3]int
	output     []int
	operations [8]func(arg uint8)
}

// init initializes the application context
func (ctx *context) init(input []string) {
	for i := range 3 {
		ctx.reg[i] = utils.Atoi(strings.Split(input[i], ": ")[1])
	}
	program := strings.Split(strings.Split(input[4], ": ")[1], ",")
	ctx.program = utils.ToUInt8Slice(program)

	// adv: division -> A
	ctx.operations[0] = func(arg uint8) {
		ctx.reg[0] = ctx.reg[0] / utils.Pow(2, ctx.combo(arg))
		ctx.increment()
	}

	// bxl: bitwise xor (B ^ op)
	ctx.operations[1] = func(arg uint8) {
		ctx.reg[1] = ctx.reg[1] ^ int(ctx.literal(arg))
		ctx.increment()
	}

	// bst: combo % 8
	ctx.operations[2] = func(arg uint8) {
		ctx.reg[1] = ctx.combo(arg) % 8
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
		ctx.output = append(ctx.output, ctx.combo(arg)%8)
		ctx.increment()
	}

	// bdv: division -> B
	ctx.operations[6] = func(arg uint8) {
		ctx.reg[1] = ctx.reg[0] / utils.Pow(2, ctx.combo(arg))
		ctx.increment()
	}

	// cdv: division -> C
	ctx.operations[7] = func(arg uint8) {
		ctx.reg[2] = ctx.reg[0] / utils.Pow(2, ctx.combo(arg))
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
	for int(ctx.ic) < len(ctx.program) {
		ctx.operations[ctx.program[ctx.ic]](ctx.program[ctx.ic+1])
	}
}

// formattedOutput formats the output slice of the program as comma-separated string
func (ctx *context) formattedOutput() string {
	return strings.Join(utils.ToStringSlice(ctx.output), ",")
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
	return ""
}
