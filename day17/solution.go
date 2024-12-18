package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	op  int
	arg int64
}

type Day17 struct {
	ra, rb, rc   int64
	program      string
	instructions []instruction
}

func (d *Day17) combo(x int64) int64 {
	switch {
	case x == 4:
		return d.ra
	case x == 5:
		return d.rb
	case x == 6:
		return d.rc
	default:
		return x
	}
}

func getInstructions(instructionString string) []instruction {
	instructions := make([]instruction, 0)
	literals := strings.Split(instructionString, ",")
	for i := 0; i < len(literals)-1; i += 2 {
		instr, _ := strconv.Atoi(literals[i])
		param, _ := strconv.ParseInt(literals[i+1], 10, 64)
		instructions = append(instructions, instruction{instr, param})
	}

	return instructions
}

func (d *Day17) Init(lines []string) {
	line := strings.Split(lines[0], ": ")
	d.ra, _ = strconv.ParseInt(line[1], 10, 64)
	line = strings.Split(lines[1], ": ")
	d.rb, _ = strconv.ParseInt(line[1], 10, 64)
	line = strings.Split(lines[2], ": ")
	d.rc, _ = strconv.ParseInt(line[1], 10, 64)

	line = strings.Split(lines[4], ": ")
	d.program = line[1]
	d.instructions = getInstructions(line[1])
}

func (d *Day17) run() string {
	ptr := 0
	builder := strings.Builder{}
	for ptr < len(d.instructions) {
		instr := d.instructions[ptr]
		ptr += 1
		switch instr.op {
		case 0:
			d.ra = d.ra / int64(math.Pow(2.0, float64(d.combo(instr.arg))))
		case 1:
			d.rb = d.rb ^ instr.arg
		case 2:
			d.rb = d.combo(instr.arg) % 8
		case 3:
			if d.ra != 0 {
				ptr = int(instr.arg)
			}
		case 4:
			d.rb = d.rb ^ d.rc
		case 5:
			builder.Write([]byte(fmt.Sprintf("%d,", d.combo(instr.arg)%8)))
		case 6:
			d.rb = d.ra / int64(math.Pow(2.0, float64(d.combo(instr.arg))))
		case 7:
			d.rc = d.ra / int64(math.Pow(2.0, float64(d.combo(instr.arg))))
		}
	}

	out := builder.String()
	if len(out) > 0 {
		out = out[:len(out)-1]
	}

	return out
}

func (d *Day17) SolveSimple() string {
	return d.run()
}

func (d *Day17) SolveAdvanced() string {
	//sum := int64(0)
	//strs := strings.Split(d.program, ",")
	//slices.Reverse(strs)
	//for _, s := range strs {
	//	v, _ := strconv.Atoi(s)
	//	sum += int64(v)
	//	sum = sum << 3
	//}
	//
	//fmt.Println(d.program)
	//for i := sum; i < sum+8; i++ {
	//	d.ra = i
	//	d.rb = 0
	//	d.rc = 0
	//	d.instructions = getInstructions(d.program)
	//	fmt.Println(d.run())
	//}
	//
	//return fmt.Sprintf("%d\n", sum)
	return ""
}
