package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	correct   [][]int
	incorrect [][]int
	precedes  map[int]map[int]bool
}

func (d *Day5) Init(lines []string) {
	d.precedes = make(map[int]map[int]bool)
	updates := make([][]int, 0)

	i := 0
	for len(lines[i]) > 0 {
		s := strings.Split(lines[i], "|")
		left, _ := strconv.Atoi(s[0])
		right, _ := strconv.Atoi(s[1])
		if d.precedes[right] == nil {
			d.precedes[right] = make(map[int]bool)
		}
		d.precedes[right][left] = true
		i++
	}

	i++
	for i < len(lines) {
		update := make([]int, 0)
		for _, s := range strings.Split(lines[i], ",") {
			v, _ := strconv.Atoi(s)
			update = append(update, v)
		}
		slices.Reverse(update)
		updates = append(updates, update)
		i++
	}

	d.correct, d.incorrect = d.classify(updates)
}

func (d *Day5) classify(updates [][]int) (correct [][]int, incorrect [][]int) {
	correct = make([][]int, 0)
	incorrect = make([][]int, 0)
	for _, update := range updates {
		ok := true
		for i, v := range update {
			for j := 0; j < i; j++ {
				if d.precedes[v][update[j]] {
					ok = false
					break
				}
			}

			if !ok {
				incorrect = append(incorrect, update)
				break
			}
		}

		if ok {
			correct = append(correct, update)
		}
	}

	return correct, incorrect
}

func (d *Day5) SolveSimple() string {
	sum := 0
	for _, correct := range d.correct {
		sum += correct[len(correct)/2]
	}

	return fmt.Sprintf("%d", sum)
}

func (d *Day5) SolveAdvanced() string {
	sum := 0
	for _, incorrect := range d.incorrect {
		ok := false
		for !ok {
			ok = true
			for i := 1; i < len(incorrect); i++ {
				if d.precedes[incorrect[i-1]][incorrect[i]] {
					ok = false
					incorrect[i-1], incorrect[i] = incorrect[i], incorrect[i-1]
				}
			}
		}
	}

	for _, incorrect := range d.incorrect {
		sum += incorrect[len(incorrect)/2]
	}

	return fmt.Sprintf("%d", sum)
}
