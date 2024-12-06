package day6

import (
	"fmt"
	"log"
)

const (
	up = iota
	left
	down
	right
)

type triple struct {
	x, y, z int
}

type Day6 struct {
	area               [][]int32
	xpos, ypos         int
	initXPos, initYPos int
	width, height      int
	direction          int
	initDirection      int
}

func (d *Day6) Init(lines []string) {
	d.width = len(lines[0])
	d.height = len(lines)
	d.area = make([][]int32, len(lines))
	for i, line := range lines {
		d.area[i] = make([]int32, len(line))
		for j, ch := range line {
			switch ch {
			case '^':
				d.direction = up
				d.xpos, d.ypos = j, i
			case '>':
				d.direction = left
				d.xpos, d.ypos = j, i
			case 'v':
				d.direction = down
				d.xpos, d.ypos = j, i
			case '<':
				d.direction = right
				d.xpos, d.ypos = j, i
			default:
				d.area[i][j] = ch
			}
		}
	}

	d.initXPos, d.initYPos = d.xpos, d.ypos
	d.initDirection = d.direction
}

func (d *Day6) SolveSimple() string {
	oob := false
	xstep, ystep := 0, 0
	for !oob {
		xstep, ystep = 0, 0
		switch d.direction {
		case up:
			if d.ypos == 0 {
				oob = true
			} else {
				ystep--
			}
		case right:
			if d.xpos == d.width-1 {
				oob = true
			} else {
				xstep++
			}
		case down:
			if d.ypos == d.height-1 {
				oob = true
			} else {
				ystep++
			}
		case left:
			if d.xpos == 0 {
				oob = true
			} else {
				xstep--
			}
		}

		if oob {
			break
		}

		if d.area[d.ypos+ystep][d.xpos+xstep] == '#' {
			d.direction = nextDirection(d.direction)
		} else {
			d.area[d.ypos][d.xpos] = 'X'
			d.xpos += xstep
			d.ypos += ystep
		}
	}

	d.area[d.ypos][d.xpos] = 'X'

	sum := 0
	for i, line := range d.area {
		for j := range line {
			if d.area[i][j] == 'X' {
				sum++
			}
		}
	}

	return fmt.Sprintf("%d", sum)
}

func nextDirection(dir int) int {
	switch dir {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		log.Fatal("invalid direction")
	}

	return up
}

func (d *Day6) SolveAdvanced() string {
	possibilities := 0
	for i, lines := range d.area {
		for j := range lines {
			if d.area[i][j] == 'X' && (i != d.initYPos || j != d.initXPos) {
				d.area[i][j] = '#'
				if d.hasLoop() {
					possibilities++
				}
				d.area[i][j] = 'X'
			}
		}
	}

	return fmt.Sprintf("%d", possibilities)
}

func (d *Day6) hasLoop() bool {
	turns := make(map[triple]bool)

	d.xpos, d.ypos = d.initXPos, d.initYPos
	d.direction = d.initDirection
	xstep, ystep := 0, 0
	for {
		xstep, ystep = 0, 0
		switch d.direction {
		case up:
			if d.ypos == 0 {
				return false
			} else {
				ystep--
			}
		case right:
			if d.xpos == d.width-1 {
				return false
			} else {
				xstep++
			}
		case down:
			if d.ypos == d.height-1 {
				return false
			} else {
				ystep++
			}
		case left:
			if d.xpos == 0 {
				return false
			} else {
				xstep--
			}
		}

		if d.area[d.ypos+ystep][d.xpos+xstep] == '#' {
			t := triple{d.xpos, d.ypos, d.direction}
			if turns[t] {
				return true
			}

			turns[t] = true
			d.direction = nextDirection(d.direction)
		} else {
			d.xpos += xstep
			d.ypos += ystep
		}
	}
}
