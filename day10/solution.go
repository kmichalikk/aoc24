package day10

import (
	"fmt"
)

type vector struct {
	x, y int
}

type Day10 struct {
	terrain       [][]int
	trailheads    []vector
	width, height int
}

func (d *Day10) canHike(from, to vector) bool {
	diff := d.terrain[to.y][to.x] - d.terrain[from.y][from.x]
	return diff == 1
}

func (d *Day10) Init(lines []string) {
	d.terrain = make([][]int, len(lines))
	d.trailheads = make([]vector, 0)
	d.width = len(lines[0])
	d.height = len(lines)
	for y, line := range lines {
		d.terrain[y] = make([]int, len(line))
		for x, ch := range line {
			d.terrain[y][x] = int(ch) - int('0')
			if d.terrain[y][x] == 0 {
				d.trailheads = append(d.trailheads, vector{x, y})
			}
		}
	}
}

func (d *Day10) SolveSimple() string {
	var visit func(vector) int
	visit = func(u vector) int {
		visited := make(map[vector]bool)
		score := 0
		q := []vector{u}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			if d.terrain[v.y][v.x] == 9 {
				score++
			}

			n, e, s, w := vector{v.x, v.y - 1}, vector{v.x + 1, v.y}, vector{v.x, v.y + 1}, vector{v.x - 1, v.y}
			if n.y >= 0 && !visited[n] && d.canHike(v, n) {
				visited[n] = true
				q = append(q, n)
			}
			if e.x <= d.width-1 && !visited[e] && d.canHike(v, e) {
				visited[e] = true
				q = append(q, e)
			}
			if s.y <= d.height-1 && !visited[s] && d.canHike(v, s) {
				visited[s] = true
				q = append(q, s)
			}
			if w.x >= 0 && !visited[w] && d.canHike(v, w) {
				visited[w] = true
				q = append(q, w)
			}
		}

		return score
	}

	total := 0
	for _, h := range d.trailheads {
		total += visit(h)
	}

	return fmt.Sprintf("%d\n", total)
}

func (d *Day10) SolveAdvanced() string {
	var visited map[vector]bool
	var totalScore map[vector]int

	var visit func(vector)
	visit = func(u vector) {
		if d.terrain[u.y][u.x] == 9 {
			totalScore[u] = 1
			return
		}

		n, e, s, w := vector{u.x, u.y - 1}, vector{u.x + 1, u.y}, vector{u.x, u.y + 1}, vector{u.x - 1, u.y}
		score := 0
		if n.y >= 0 && d.canHike(u, n) {
			if !visited[n] {
				visit(n)
			}
			score += totalScore[n]
		}
		if e.x < d.width && d.canHike(u, e) {
			if !visited[e] {
				visit(e)
			}
			score += totalScore[e]
		}
		if s.y < d.height && d.canHike(u, s) {
			if !visited[s] {
				visit(s)
			}
			score += totalScore[s]
		}
		if w.x >= 0 && d.canHike(u, w) {
			if !visited[w] {
				visit(w)
			}
			score += totalScore[w]
		}

		totalScore[u] = score
		visited[u] = true
	}

	total := 0
	for _, h := range d.trailheads {
		visited = make(map[vector]bool)
		totalScore = make(map[vector]int)
		visit(h)
		total += totalScore[h]
	}

	return fmt.Sprintf("%d\n", total)
}
