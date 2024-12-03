package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

type Day3 struct {
	lines []string
}

func (d *Day3) Init(lines []string) {
	d.lines = lines
}

func (d *Day3) SolveSimple() string {
	re, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	var total int64
	for _, line := range d.lines {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			total += int64(a * b)
		}
	}

	return fmt.Sprint(total)
}

func (d *Day3) SolveAdvanced() string {
	return ""
}
