package day14

import (
	"fmt"
	"regexp"
	"strconv"
)

type vector struct {
	x, y int
}

type robot struct {
	position, velocity vector
}

type Day14 struct {
	robots []robot
	box    vector
}

func (v vector) addModulo(other, modulo vector) vector {
	xm, ym := (v.x+other.x)%modulo.x, (v.y+other.y)%modulo.y
	if xm < 0 {
		xm = modulo.x + xm
	}
	if ym < 0 {
		ym = modulo.y + ym
	}
	return vector{xm, ym}
}

func (d *Day14) Init(lines []string) {
	d.box.x = 101
	d.box.y = 103
	d.robots = make([]robot, 0)
	lineRegex, _ := regexp.Compile("p=(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")
	for _, line := range lines {
		match := lineRegex.FindStringSubmatch(line)
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		d.robots = append(d.robots, robot{vector{px, py}, vector{vx, vy}})
	}
}

func (d *Day14) SolveSimple() string {
	safety := 1
	positions := make(map[vector]int)
	for _, robot := range d.robots {
		pos := robot.position.addModulo(vector{100 * robot.velocity.x, 100 * robot.velocity.y}, d.box)
		positions[pos] = positions[pos] + 1
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	xHalf, yHalf := d.box.x/2, d.box.y/2
	for pos, count := range positions {
		if pos.x < xHalf && pos.y < yHalf {
			q1 += count
		} else if pos.x > xHalf && pos.y < yHalf {
			q2 += count
		} else if pos.x < xHalf && pos.y > yHalf {
			q3 += count
		} else if pos.x > xHalf && pos.y > yHalf {
			q4 += count
		}
	}

	safety = q1 * q2 * q3 * q4

	return fmt.Sprintf("%d\n", safety)
}

func (d *Day14) SolveAdvanced() string {
	for i := range 10000 {
		positions := make(map[vector]int)
		for _, robot := range d.robots {
			pos := robot.position.addModulo(vector{i * robot.velocity.x, i * robot.velocity.y}, d.box)
			positions[pos] = positions[pos] + 1
		}

		q1, q2, q3, q4 := 0, 0, 0, 0
		xHalf, yHalf := d.box.x/2, d.box.y/2
		for pos, count := range positions {
			if pos.x < xHalf && pos.y < yHalf {
				q1 += count
			} else if pos.x > xHalf && pos.y < yHalf {
				q2 += count
			} else if pos.x < xHalf && pos.y > yHalf {
				q3 += count
			} else if pos.x > xHalf && pos.y > yHalf {
				q4 += count
			}
		}

		if q1 > len(d.robots)/2 || q2 > len(d.robots)/2 || q3 > len(d.robots)/2 || q4 > len(d.robots)/2 {
			for j := 0; j < d.box.y; j++ {
				for k := 0; k < d.box.x; k++ {
					if positions[vector{k, j}] > 0 {
						fmt.Printf("x")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Printf("\n")
			}

			return fmt.Sprintf("%d\n", i)
		}
	}

	return ""
}
