package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type equation struct {
	value    int64
	operands []int64
}

type Day7 struct {
	equations []equation
}

func (d *Day7) Init(lines []string) {
	d.equations = make([]equation, len(lines))
	for i, line := range lines {
		tok := strings.Split(line, " ")
		value, _ := strconv.ParseInt(tok[0][:len(tok[0])-1], 10, 64)
		ops := make([]int64, len(tok)-1)
		for i, op := range tok[1:] {
			ops[i], _ = strconv.ParseInt(op, 10, 64)
		}

		d.equations[i] = equation{value, ops}
	}
}

func (d *Day7) SolveSimple() string {
	var tryMatchValue func(int, int, int64) bool
	tryMatchValue = func(index, depth int, total int64) bool {
		operands := d.equations[index].operands
		if depth < len(operands) {
			return tryMatchValue(index, depth+1, total+operands[depth]) ||
				tryMatchValue(index, depth+1, total*operands[depth])
		}

		return total == d.equations[index].value
	}

	var sum int64 = 0
	for i := range d.equations {
		if tryMatchValue(i, 1, d.equations[i].operands[0]) {
			sum += d.equations[i].value
		}
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day7) SolveAdvanced() string {
	var tryMatchValue func(int, int, int64) bool
	tryMatchValue = func(index, depth int, total int64) bool {
		operands := d.equations[index].operands
		if depth < len(operands) {
			concat, _ := strconv.ParseInt(fmt.Sprintf("%d%d", total, operands[depth]), 10, 64)
			return tryMatchValue(index, depth+1, total+operands[depth]) ||
				tryMatchValue(index, depth+1, total*operands[depth]) ||
				tryMatchValue(index, depth+1, concat)
		}

		return total == d.equations[index].value
	}

	var sum int64 = 0
	for i := range d.equations {
		if tryMatchValue(i, 1, d.equations[i].operands[0]) {
			sum += d.equations[i].value
		}
	}

	return fmt.Sprintf("%d", sum)
}
