package day15

import (
	"fmt"
)

type vector struct {
	x, y int
}

const (
	empty = iota
	wall
	boxes
	up
	down
	left
	right
	leftBox
	rightBox
)

type Day15 struct {
	initialPosition vector
	boxes           []vector
	walls           []vector
	width, height   int
	moves           []int
}

func (d *Day15) Init(lines []string) {
	d.boxes = make([]vector, 0)
	d.walls = make([]vector, 0)
	d.moves = make([]int, 0)
	i := 0
	for len(lines[i]) > 0 {
		for j, ch := range lines[i] {
			if ch == '#' {
				d.walls = append(d.walls, vector{j, i})
			} else if ch == 'O' {
				d.boxes = append(d.boxes, vector{j, i})
			} else if ch == '@' {
				d.initialPosition = vector{j, i}
			}
		}
		i++
	}

	d.height = i
	d.width = len(lines[0])

	for ; i < len(lines); i++ {
		for _, ch := range lines[i] {
			switch ch {
			case '^':
				d.moves = append(d.moves, up)
			case '>':
				d.moves = append(d.moves, right)
			case 'v':
				d.moves = append(d.moves, down)
			case '<':
				d.moves = append(d.moves, left)
			}
		}
	}
}

func (d *Day15) SolveSimple() string {
	pos := d.initialPosition
	area := make([][]int, d.height)
	for i := range area {
		area[i] = make([]int, d.width)
		for j := range area[i] {
			area[i][j] = empty
		}
	}

	for _, v := range d.boxes {
		area[v.y][v.x] = boxes
	}

	for _, v := range d.walls {
		area[v.y][v.x] = wall
	}

	for _, m := range d.moves {
		switch m {
		case up:
			y := pos.y - 1
			for area[y][pos.x] == boxes {
				y--
			}

			if area[y][pos.x] == empty {
				y++
				for y < pos.y {
					area[y][pos.x], area[y-1][pos.x] = area[y-1][pos.x], area[y][pos.x]
					y++
				}
				pos.y--
			}
		case down:
			y := pos.y + 1
			for area[y][pos.x] == boxes {
				y++
			}

			if area[y][pos.x] == empty {
				y--
				for y > pos.y {
					area[y][pos.x], area[y+1][pos.x] = area[y+1][pos.x], area[y][pos.x]
					y--
				}
				pos.y++
			}
		case right:
			x := pos.x + 1
			for area[pos.y][x] == boxes {
				x++
			}

			if area[pos.y][x] == empty {
				x--
				for x > pos.x {
					area[pos.y][x], area[pos.y][x+1] = area[pos.y][x+1], area[pos.y][x]
					x--
				}
				pos.x++
			}
		case left:
			x := pos.x - 1
			for area[pos.y][x] == boxes {
				x--
			}

			if area[pos.y][x] == empty {
				x++
				for x < pos.x {
					area[pos.y][x], area[pos.y][x-1] = area[pos.y][x-1], area[pos.y][x]
					x++
				}
				pos.x--
			}
		}
	}

	total := 0
	for i := range area {
		for j := range area[i] {
			if area[i][j] == boxes {
				total += i*100 + j
			}
		}
	}

	return fmt.Sprintf("%d\n", total)
}

func canMove(area *[][]int, boxPos vector, direction int) bool {
	if (*area)[boxPos.y][boxPos.x] == wall {
		return false
	}

	if (*area)[boxPos.y][boxPos.x] == rightBox {
		boxPos.x--
	}

	switch direction {
	case up:
		return ((*area)[boxPos.y-1][boxPos.x] == empty || canMove(area, vector{boxPos.x, boxPos.y - 1}, direction)) &&
			((*area)[boxPos.y-1][boxPos.x+1] == empty || canMove(area, vector{boxPos.x + 1, boxPos.y - 1}, direction))
	case down:
		return ((*area)[boxPos.y+1][boxPos.x] == empty || canMove(area, vector{boxPos.x, boxPos.y + 1}, direction)) &&
			((*area)[boxPos.y+1][boxPos.x+1] == empty || canMove(area, vector{boxPos.x + 1, boxPos.y + 1}, direction))
	case right:
		return (*area)[boxPos.y][boxPos.x+2] == empty || canMove(area, vector{boxPos.x + 2, boxPos.y}, direction)
	case left:
		return (*area)[boxPos.y][boxPos.x-1] == empty || canMove(area, vector{boxPos.x - 2, boxPos.y}, direction)
	}

	return false
}

func doMove(area *[][]int, boxPos vector, direction int) {
	// there are two cases:
	// a) box is directly above or
	// b) there are two boxes side-by-side
	// for a), the second attempt will be to move empty space, so we need to return immediately
	if (*area)[boxPos.y][boxPos.x] == empty {
		return
	}

	if (*area)[boxPos.y][boxPos.x] == rightBox {
		boxPos.x--
	}

	switch direction {
	case up:
		doMove(area, vector{boxPos.x, boxPos.y - 1}, direction)
		doMove(area, vector{boxPos.x + 1, boxPos.y - 1}, direction) // returns immediately if a) occurs
		(*area)[boxPos.y][boxPos.x], (*area)[boxPos.y-1][boxPos.x] = (*area)[boxPos.y-1][boxPos.x], (*area)[boxPos.y][boxPos.x]
		(*area)[boxPos.y][boxPos.x+1], (*area)[boxPos.y-1][boxPos.x+1] = (*area)[boxPos.y-1][boxPos.x+1], (*area)[boxPos.y][boxPos.x+1]
	case down:
		doMove(area, vector{boxPos.x, boxPos.y + 1}, direction)
		doMove(area, vector{boxPos.x + 1, boxPos.y + 1}, direction) // returns immediately if a) occurs
		(*area)[boxPos.y][boxPos.x], (*area)[boxPos.y+1][boxPos.x] = (*area)[boxPos.y+1][boxPos.x], (*area)[boxPos.y][boxPos.x]
		(*area)[boxPos.y][boxPos.x+1], (*area)[boxPos.y+1][boxPos.x+1] = (*area)[boxPos.y+1][boxPos.x+1], (*area)[boxPos.y][boxPos.x+1]
	case right:
		doMove(area, vector{boxPos.x + 2, boxPos.y}, direction)
		(*area)[boxPos.y][boxPos.x+1], (*area)[boxPos.y][boxPos.x+2] = (*area)[boxPos.y][boxPos.x+2], (*area)[boxPos.y][boxPos.x+1]
		(*area)[boxPos.y][boxPos.x], (*area)[boxPos.y][boxPos.x+1] = (*area)[boxPos.y][boxPos.x+1], (*area)[boxPos.y][boxPos.x]
	case left:
		doMove(area, vector{boxPos.x - 1, boxPos.y}, direction)
		(*area)[boxPos.y][boxPos.x], (*area)[boxPos.y][boxPos.x-1] = (*area)[boxPos.y][boxPos.x-1], (*area)[boxPos.y][boxPos.x]
		(*area)[boxPos.y][boxPos.x+1], (*area)[boxPos.y][boxPos.x] = (*area)[boxPos.y][boxPos.x], (*area)[boxPos.y][boxPos.x+1]
	}
}

func (d *Day15) SolveAdvanced() string {
	pos := vector{d.initialPosition.x * 2, d.initialPosition.y}
	area := make([][]int, d.height)
	for i := range area {
		area[i] = make([]int, d.width*2)
		for j := range area[i] {
			area[i][j] = empty
		}
	}

	for _, v := range d.boxes {
		area[v.y][v.x*2] = leftBox
		area[v.y][v.x*2+1] = rightBox
	}

	for _, v := range d.walls {
		area[v.y][v.x*2] = wall
		area[v.y][v.x*2+1] = wall
	}

	for _, m := range d.moves {
		switch m {
		case up:
			if area[pos.y-1][pos.x] == empty {
				pos.y--
			} else if canMove(&area, vector{pos.x, pos.y - 1}, m) {
				doMove(&area, vector{pos.x, pos.y - 1}, m)
				pos.y--
			}
		case down:
			if area[pos.y+1][pos.x] == empty {
				pos.y++
			} else if canMove(&area, vector{pos.x, pos.y + 1}, m) {
				doMove(&area, vector{pos.x, pos.y + 1}, m)
				pos.y++
			}
		case right:
			if area[pos.y][pos.x+1] == empty {
				pos.x++
			} else if canMove(&area, vector{pos.x + 1, pos.y}, m) {
				doMove(&area, vector{pos.x + 1, pos.y}, m)
				pos.x++
			}
		case left:
			if area[pos.y][pos.x-1] == empty {
				pos.x--
			} else if canMove(&area, vector{pos.x - 1, pos.y}, m) {
				doMove(&area, vector{pos.x - 1, pos.y}, m)
				pos.x--
			}
		}
	}

	total := 0
	for i := range area {
		for j := range area[i] {
			if area[i][j] == leftBox {
				total += i*100 + j
			}
		}
	}

	return fmt.Sprintf("%d", total)
}
