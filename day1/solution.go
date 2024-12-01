package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day1 struct {
	lines []string
	left  []int64
	right []int64
}

func (d *Day1) Init(lines []string) {
	d.lines = lines
	d.left = make([]int64, 0, len(d.lines))
	d.right = make([]int64, 0, len(d.lines))

	for _, line := range d.lines {
		tok := strings.SplitN(line, "   ", 2)
		l, _ := strconv.Atoi(tok[0])
		r, _ := strconv.Atoi(tok[1])

		d.left = append(d.left, int64(l))
		d.right = append(d.right, int64(r))

		slices.Sort(d.left)
		slices.Sort(d.right)
	}
}

func (d *Day1) SolveSimple() string {

	var sum int64

	for i := range len(d.left) {
		if diff := d.right[i] - d.left[i]; diff > 0 {
			sum += diff
		} else {
			sum -= diff
		}
	}

	return fmt.Sprint(sum)
}

func (d *Day1) SolveAdvanced() string {
	counts := make(map[int64]int)
	currentCount := 0
	currentValue := d.right[0]
	for i := range len(d.right) {
		if d.right[i] == currentValue {
			currentCount++
		} else {
			counts[currentValue] = currentCount
			currentValue = d.right[i]
			currentCount = 1
		}
	}

	counts[currentValue] = currentCount

	var sum int64

	for i := range d.left {
		sum += d.left[i] * int64(counts[d.left[i]])
	}

	return fmt.Sprint(sum)
}
