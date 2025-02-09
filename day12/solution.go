package day12

import "fmt"

type Day12 struct {
	area          [][]int32
	width, height int
}

type region struct {
	area, metric int
}

type vector struct {
	x, y int
}

func (d *Day12) Init(lines []string) {
	d.area = make([][]int32, len(lines)+2)
	d.width = len(lines[0])
	d.height = len(lines)
	d.area[0] = make([]int32, len(lines[0])+2)
	d.area[len(d.area)-1] = make([]int32, len(lines[0])+2)
	for i, line := range lines {
		d.area[i+1] = make([]int32, len(line)+2)
		for j, ch := range line {
			d.area[i+1][j+1] = ch
		}
	}
}

func (d *Day12) SolveSimple() string {
	regions := make([]region, 0)
	visited := make(map[vector]bool)

	for i := 1; i < d.height+1; i++ {
		for j := 1; j < d.width+1; j++ {
			v := vector{j, i}
			if visited[v] {
				continue
			}

			q := []vector{v}
			r := region{}
			for len(q) > 0 {
				u := q[0]
				q = q[1:]
				if visited[u] {
					continue
				}

				r.area++
				visited[u] = true

				n := vector{u.x, u.y - 1}
				if d.area[u.y][u.x] != d.area[u.y-1][u.x] {
					r.metric++
				} else {
					q = append(q, n)
				}

				s := vector{u.x, u.y + 1}
				if d.area[u.y][u.x] != d.area[u.y+1][u.x] {
					r.metric++
				} else {
					q = append(q, s)
				}

				w := vector{u.x - 1, u.y}
				if d.area[u.y][u.x] != d.area[u.y][u.x-1] {
					r.metric++
				} else {
					q = append(q, w)
				}

				e := vector{u.x + 1, u.y}
				if d.area[u.y][u.x] != d.area[u.y][u.x+1] {
					r.metric++
				} else {
					q = append(q, e)
				}
			}

			regions = append(regions, r)
		}
	}

	total := 0
	for _, r := range regions {
		total += r.area * r.metric
	}

	return fmt.Sprintf("%d\n", total)
}

func (d *Day12) SolveAdvanced() string {
	regions := make([]region, 0)
	visited := make(map[vector]bool)

	for i := 1; i < d.height+1; i++ {
		for j := 1; j < d.width+1; j++ {
			v := vector{j, i}
			if visited[v] {
				continue
			}

			q := []vector{v}
			r := region{}
			for len(q) > 0 {
				u := q[0]
				q = q[1:]
				if visited[u] {
					continue
				}

				r.area++
				visited[u] = true
				outsideCount := 0

				n := vector{u.x, u.y - 1}
				if d.area[u.y][u.x] != d.area[u.y-1][u.x] {
					outsideCount++
				} else {
					q = append(q, n)
				}

				s := vector{u.x, u.y + 1}
				if d.area[u.y][u.x] != d.area[u.y+1][u.x] {
					outsideCount++
				} else {
					q = append(q, s)
				}

				w := vector{u.x - 1, u.y}
				if d.area[u.y][u.x] != d.area[u.y][u.x-1] {
					outsideCount++
				} else {
					q = append(q, w)
				}

				e := vector{u.x + 1, u.y}
				if d.area[u.y][u.x] != d.area[u.y][u.x+1] {
					outsideCount++
				} else {
					q = append(q, e)
				}

				below := d.area[u.y][u.x]
				switch outsideCount {
				case 4:
					r.metric += 8
				case 3:
					r.metric += 4
				case 2:
					if (d.area[u.y][u.x+1] != below && d.area[u.y+1][u.x] != below) ||
						(d.area[u.y][u.x-1] != below && d.area[u.y+1][u.x] != below) ||
						(d.area[u.y][u.x-1] != below && d.area[u.y-1][u.x] != below) ||
						(d.area[u.y][u.x+1] != below && d.area[u.y-1][u.x] != below) {
						r.metric += 2
					}
				}
				if d.area[u.y][u.x+1] == below && d.area[u.y+1][u.x] == below && d.area[u.y+1][u.x+1] != below {
					r.metric += 2
				}
				if d.area[u.y][u.x-1] == below && d.area[u.y+1][u.x] == below && d.area[u.y+1][u.x-1] != below {
					r.metric += 2
				}
				if d.area[u.y][u.x-1] == below && d.area[u.y-1][u.x] == below && d.area[u.y-1][u.x-1] != below {
					r.metric += 2
				}
				if d.area[u.y][u.x+1] == below && d.area[u.y-1][u.x] == below && d.area[u.y-1][u.x+1] != below {
					r.metric += 2
				}
			}

			r.metric /= 2
			regions = append(regions, r)
		}
	}

	total := 0
	for _, r := range regions {
		total += r.area * r.metric
	}

	return fmt.Sprintf("%d\n", total)
}
