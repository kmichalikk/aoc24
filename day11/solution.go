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

		if p.v == 0 {
			stack = append(stack, pair{p.n + 1, 1})
			continue
		}

		w := int(math.Log10(float64(p.v))) + 1
		if w%2 == 0 {
			div := int64(math.Pow10(w / 2))
			stack = append(stack, pair{p.n + 1, p.v / div})
			stack = append(stack, pair{p.n + 1, p.v % div})
			continue
		}

		stack = append(stack, pair{p.n + 1, 2024 * p.v})
	}

	return fmt.Sprintf("%d\n", total)
}

func (d *Day11) SolveAdvanced() string {
	d.steps = 75
	var total int64 = 0
	memo := make(map[pair]int64)
	var visit func(int64, int) int64
	visit = func(v int64, depth int) int64 {
		p := pair{depth, v}
		m, ok := memo[p]
		if ok {
			return m
		}

		if depth == d.steps {
			return 1
		}

		if p.v == 0 {
			memo[p] = visit(1, depth+1)
			return memo[p]
		}

		w := int(math.Log10(float64(p.v))) + 1
		if w%2 == 0 {
			div := int64(math.Pow10(w / 2))
			memo[p] = visit(p.v/div, depth+1) + visit(p.v%div, depth+1)
			return memo[p]
		}

		memo[p] = visit(2024*p.v, depth+1)
		return memo[p]
	}

	for i := range d.numbers {
		total += visit(d.numbers[len(d.numbers)-i-1], 0)
	}

	return fmt.Sprintf("%d\n", total)
}
