package day19

import (
	"fmt"
	"strings"
)

type Day19 struct {
	towels   []string
	patterns []string
}

func (d *Day19) Init(lines []string) {
	d.towels = strings.Split(lines[0], ", ")

	for _, line := range lines[2:] {
		d.patterns = append(d.patterns, line)
	}
}

func (d *Day19) SolveSimple() string {
	var testPattern func(string, *map[int]bool) bool
	testPattern = func(pattern string, dynamicTable *map[int]bool) bool {
		patternLength := len(pattern)
		if patternLength == 0 || (*dynamicTable)[patternLength] {
			return true
		}

		for _, towel := range d.towels {
			towelLength := len(towel)
			if towelLength > patternLength || towel != pattern[:towelLength] {
				continue
			}

			if testPattern(pattern[towelLength:], dynamicTable) {
				(*dynamicTable)[patternLength] = true
				return true
			}
		}

		return false
	}

	possible := 0
	for _, pattern := range d.patterns {
		dynamicTable := make(map[int]bool)
		if testPattern(pattern, &dynamicTable) {
			possible++
		}
	}

	return fmt.Sprintf("%d", possible)
}

func (d *Day19) SolveAdvanced() string {
	var countPattern func(string, *map[int]int) int
	countPattern = func(pattern string, dynamicTable *map[int]int) int {
		patternLength := len(pattern)
		if count, ok := (*dynamicTable)[patternLength]; ok {
			return count
		}

		if patternLength == 0 {
			return 1
		}

		for _, towel := range d.towels {
			towelLength := len(towel)
			if towelLength > patternLength || towel != pattern[:towelLength] {
				continue
			}

			(*dynamicTable)[patternLength] += countPattern(pattern[towelLength:], dynamicTable)
		}

		return (*dynamicTable)[patternLength]
	}

	count := 0
	for _, pattern := range d.patterns {
		dynamicTable := make(map[int]int)
		count += countPattern(pattern, &dynamicTable)
	}

	return fmt.Sprintf("%d", count)
}
