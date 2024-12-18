package day18

import (
	"fmt"
	"strconv"
	"strings"
)

type vector struct {
	x, y int
}

type item struct {
	v    vector
	wave int
}

type Day18 struct {
	lines           []string
	corruptionTimes map[vector]int
	width, height   int
}

func (d *Day18) Init(lines []string) {
	d.corruptionTimes = make(map[vector]int)
	d.lines = lines
	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		d.width = max(d.width, x+1)
		d.height = max(d.height, y+1)
		d.corruptionTimes[vector{x, y}] = i
	}

	for i := range d.height {
		for j := range d.width {
			_, ok := d.corruptionTimes[vector{j, i}]
			if !ok {
				d.corruptionTimes[vector{j, i}] = 2 * d.width * d.height // approx. infinity
			}
		}
	}
}

func (d *Day18) bfs(threshold int) (int, bool) {
	queue := []item{{vector{0, 0}, 0}}
	visited := make(map[vector]bool)
	found := false
	steps := 0
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]

		if it.wave > steps {
			steps = it.wave
		}

		if it.v.x == d.width-1 && it.v.y == d.height-1 {
			found = true
			break
		}

		if visited[it.v] || d.corruptionTimes[it.v] < it.wave {
			continue
		}

		visited[it.v] = true

		up := vector{it.v.x, it.v.y - 1}
		down := vector{it.v.x, it.v.y + 1}
		right := vector{it.v.x + 1, it.v.y}
		left := vector{it.v.x - 1, it.v.y}
		if up.y >= 0 && !visited[up] && d.corruptionTimes[up] > threshold {
			queue = append(queue, item{up, it.wave + 1})
		}
		if down.y < d.height && !visited[down] && d.corruptionTimes[down] > threshold {
			queue = append(queue, item{down, it.wave + 1})
		}
		if right.x < d.width && !visited[right] && d.corruptionTimes[right] > threshold {
			queue = append(queue, item{right, it.wave + 1})
		}
		if left.x >= 0 && !visited[left] && d.corruptionTimes[left] > threshold {
			queue = append(queue, item{left, it.wave + 1})
		}
	}

	return steps, found
}

func (d *Day18) SolveSimple() string {
	steps, _ := d.bfs(1023)
	return fmt.Sprintf("%d\n", steps)
}

func (d *Day18) SolveAdvanced() string {
	i := 1023
	for ; i < len(d.lines); i++ {
		d.Init(d.lines)
		_, ok := d.bfs(i)
		if !ok {
			break
		}
	}

	var x, y int
	for v, t := range d.corruptionTimes {
		if t == i {
			x, y = v.x, v.y
		}
	}

	return fmt.Sprintf("%d,%d\n", x, y)
}
