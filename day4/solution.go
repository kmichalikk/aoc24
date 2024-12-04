package day4

import (
	"fmt"
)

type Day4 struct {
	lines []string
}

func (d *Day4) Init(lines []string) {
	d.lines = lines
}

func (d *Day4) SolveSimple() string {
	xmasCount := 0
	count, length := len(d.lines), len(d.lines[0])
	for i := range d.lines {
		for j := range d.lines[i] {
			if j <= length-4 && xmas(d.lines[i][j], d.lines[i][j+1], d.lines[i][j+2], d.lines[i][j+3]) {
				xmasCount++
			}

			if i <= count-4 {
				if xmas(d.lines[i][j], d.lines[i+1][j], d.lines[i+2][j], d.lines[i+3][j]) {
					xmasCount++
				}

				if j >= 3 && xmas(d.lines[i][j], d.lines[i+1][j-1], d.lines[i+2][j-2], d.lines[i+3][j-3]) {
					xmasCount++
				}

				if j <= length-4 && xmas(d.lines[i][j], d.lines[i+1][j+1], d.lines[i+2][j+2], d.lines[i+3][j+3]) {
					xmasCount++
				}
			}
		}
	}

	return fmt.Sprintf("%d", xmasCount)
}

func xmas(a, b, c, d uint8) bool {
	return (a == 'X' && b == 'M' && c == 'A' && d == 'S') || (a == 'S' && b == 'A' && c == 'M' && d == 'X')
}

func (d *Day4) SolveAdvanced() string {
	xmasCount := 0
	count := len(d.lines)
	for i := 1; i < count-1; i++ {
		for j := 1; j < count-1; j++ {
			if d.lines[i][j] == 'A' {
				diag1 := d.lines[i-1][j-1] == 'M' && d.lines[i+1][j+1] == 'S' || d.lines[i-1][j-1] == 'S' && d.lines[i+1][j+1] == 'M'
				diag2 := d.lines[i-1][j+1] == 'M' && d.lines[i+1][j-1] == 'S' || d.lines[i-1][j+1] == 'S' && d.lines[i+1][j-1] == 'M'
				if diag1 && diag2 {
					xmasCount++
				}
			}
		}
	}

	return fmt.Sprintf("%d", xmasCount)
}
