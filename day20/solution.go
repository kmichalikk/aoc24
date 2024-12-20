package day20

import (
	"fmt"
)

type vector struct {
	x, y int
}

type cheat struct {
	start, end vector
}

type Day20 struct {
	area          [][]int32
	width, height int
	start, end    vector
}

func (d *Day20) Init(lines []string) {
	d.width = len(lines[0])
	d.height = len(lines)
	d.area = make([][]int32, d.height)
	for i, line := range lines {
		d.area[i] = make([]int32, d.width)
		for j, ch := range line {
			if ch == 'S' {
				d.start = vector{j, i}
			}
			if ch == 'E' {
				d.end = vector{j, i}
			}

			if ch == '#' {
				d.area[i][j] = '#'
			} else {
				d.area[i][j] = ' '
			}
		}
	}
}

func (d *Day20) traverse() (map[vector]int, map[vector]bool) {
	distance := make(map[vector]int)
	visited := make(map[vector]bool)
	queue := []vector{d.start}
	currentDistance := 0
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		if visited[u] {
			continue
		}
		visited[u] = true
		distance[u] = currentDistance
		currentDistance++

		n := vector{u.x, u.y - 1}
		if n.y >= 0 && d.area[n.y][n.x] != '#' && !visited[n] {
			queue = append(queue, n)
		}

		s := vector{u.x, u.y + 1}
		if s.y < d.height && d.area[s.y][s.x] != '#' && !visited[s] {
			queue = append(queue, s)
		}

		e := vector{u.x + 1, u.y}
		if e.x < d.width && d.area[e.y][e.x] != '#' && !visited[e] {
			queue = append(queue, e)
		}

		w := vector{u.x - 1, u.y}
		if w.x >= 0 && d.area[w.y][w.x] != '#' && !visited[w] {
			queue = append(queue, w)
		}
	}

	return distance, visited
}

func (d *Day20) SolveSimple() string {
	distance, visited := d.traverse()
	cheats := make(map[cheat]int)

	for v, _ := range visited {
		n := vector{v.x, v.y - 2}
		if dn, ok := distance[n]; ok && distance[v]+2 < dn {
			cheats[cheat{v, n}] = dn - distance[v] - 2
		}

		s := vector{v.x, v.y + 2}
		if ds, ok := distance[s]; ok && distance[v]+2 < ds {
			cheats[cheat{v, s}] = ds - distance[v] - 2
		}

		e := vector{v.x + 2, v.y}
		if de, ok := distance[e]; ok && distance[v]+2 < de {
			cheats[cheat{v, e}] = de - distance[v] - 2
		}

		w := vector{v.x - 2, v.y}
		if dw, ok := distance[w]; ok && distance[v]+2 < dw {
			cheats[cheat{v, w}] = dw - distance[v] - 2
		}
	}

	count := 0
	for _, v := range cheats {
		if v >= 100 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func iabs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func (d *Day20) SolveAdvanced() string {
	distance, visited := d.traverse()
	cheats := make(map[cheat]int)

	for v, _ := range visited {
		for y := -20; y <= 20; y++ {
			for x := iabs(y) - 20; x <= 20-iabs(y); x++ {
				u := vector{v.x - x, v.y - y}
				cheatDistance := iabs(x) + iabs(y)
				if dist, ok := distance[u]; ok && distance[v]+cheatDistance < dist {
					cheats[cheat{v, u}] = dist - distance[v] - cheatDistance
				}
			}
		}
	}

	count := 0
	for _, v := range cheats {
		if v >= 100 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}
