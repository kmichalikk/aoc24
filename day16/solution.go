package day16

import (
	"container/heap"
	"fmt"
)

type position struct {
	x, y int
}

type vector struct {
	p position
	d int
}

type Day16 struct {
	area       [][]int32
	start, end position
}

const (
	up = iota
	down
	left
	right
)

type pqItem struct {
	vec    vector
	metric int
}

type priorityQueue []*pqItem

func (p *priorityQueue) Len() int {
	return len(*p)
}

func (p *priorityQueue) Less(i, j int) bool {
	return (*p)[i].metric < (*p)[j].metric
}

func (p *priorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(*pqItem))
}

func (p *priorityQueue) Pop() any {
	n := len(*p) - 1
	it := (*p)[n]
	*p = (*p)[0:n]
	return it
}

func (d *Day16) Init(lines []string) {
	d.area = make([][]int32, len(lines))
	for i, line := range lines {
		d.area[i] = make([]int32, len(line))
		for j, ch := range line {
			if ch == 'S' {
				d.start = position{j, i}
			} else if ch == 'E' {
				d.end = position{j, i}
			} else {
				d.area[i][j] = ch
			}
		}
	}
}

func (d *Day16) SolveSimple() string {
	pq := priorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &pqItem{vector{d.start, right}, 0})
	visited := make(map[vector]bool)
	var result int
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		if visited[item.vec] {
			continue
		}

		visited[item.vec] = true
		pos := item.vec.p
		dir := item.vec.d

		if pos == d.end {
			result = item.metric
			break
		}

		// up
		if pos.y-1 >= 0 && dir != down && d.area[pos.y-1][pos.x] != '#' {
			if dir == up {
				heap.Push(&pq, &pqItem{vector{position{pos.x, pos.y - 1}, up}, item.metric + 1})
			} else {
				heap.Push(&pq, &pqItem{vector{pos, up}, item.metric + 1000})
			}
		}
		// down
		if pos.y+1 < len(d.area) && dir != up && d.area[pos.y+1][pos.x] != '#' {
			if dir == down {
				heap.Push(&pq, &pqItem{vector{position{pos.x, pos.y + 1}, down}, item.metric + 1})
			} else {
				heap.Push(&pq, &pqItem{vector{pos, down}, item.metric + 1000})
			}
		}
		// right
		if pos.x+1 < len(d.area[0]) && dir != left && d.area[pos.y][pos.x+1] != '#' {
			if dir == right {
				heap.Push(&pq, &pqItem{vector{position{pos.x + 1, pos.y}, right}, item.metric + 1})
			} else {
				heap.Push(&pq, &pqItem{vector{pos, right}, item.metric + 1000})
			}
		}
		// left
		if pos.x-1 >= 0 && dir != right && d.area[pos.y][pos.x-1] != '#' {
			if dir == left {
				heap.Push(&pq, &pqItem{vector{position{pos.x - 1, pos.y}, left}, item.metric + 1})
			} else {
				heap.Push(&pq, &pqItem{vector{pos, left}, item.metric + 1000})
			}
		}
	}

	return fmt.Sprintf("%d\n", result)
}

func (d *Day16) SolveAdvanced() string {
	return ""
}
