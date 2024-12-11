package day11

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pair struct {
	n int
	v int64
}

type Day11 struct {
	numbers []int64
	steps   int
}

func (d *Day11) Init(lines []string) {
	d.numbers = make([]int64, 0)
	d.steps = 25
	for _, s := range strings.Split(lines[0], " ") {
		v, _ := strconv.Atoi(s)
		d.numbers = append(d.numbers, int64(v))
	}
}

func (d *Day11) SolveSimple() string {
	stack := make([]pair, len(d.numbers))
	for i, n := range d.numbers {
		stack[len(d.numbers)-i-1] = pair{0, n}
	}

	total := 0
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.n == d.steps {
			total += 1
			continue
		}

		w := int(math.Log10(float64(p.v))) + 1

		switch {
		case p.v == 0:
			stack = append(stack, pair{p.n + 1, 1})
		case w%2 == 0:
			div := int64(math.Pow10(w / 2))
			stack = append(stack, pair{p.n + 1, p.v / div})
			stack = append(stack, pair{p.n + 1, p.v % div})
		default:
			stack = append(stack, pair{p.n + 1, 2024 * p.v})
		}
	}

	return fmt.Sprintf("%d\n", total)
}

func (d *Day11) SolveAdvanced() string {
	d.steps = 45
	return d.SolveSimple()
}
