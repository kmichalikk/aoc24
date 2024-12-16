package day16

import (
	"container/heap"
	"fmt"
	"slices"
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

func (d *Day16) dijkstra(initial []pqItem) [][][]int {
	distances := make([][][]int, len(d.area))
	for i := range distances {
		distances[i] = make([][]int, len(d.area[i]))
		for j := range distances[i] {
			val := 2000 * len(d.area) * len(d.area[0])
			distances[i][j] = []int{val, val, val, val}
		}
	}

	pq := priorityQueue{}
	heap.Init(&pq)
	for _, it := range initial {
		heap.Push(&pq, &it)
	}

	visited := make(map[vector]bool)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		if visited[item.vec] {
			continue
		}

		visited[item.vec] = true
		pos := item.vec.p
		dir := item.vec.d
		distances[pos.y][pos.x][dir] = item.metric

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

	return distances
}

func (d *Day16) SolveSimple() string {
	distances := d.dijkstra([]pqItem{{vector{d.start, right}, 0}})

	return fmt.Sprintf("%d\n", slices.Min(distances[d.end.y][d.end.x]))
}

func (d *Day16) SolveAdvanced() string {
	distancesFromStart := d.dijkstra([]pqItem{{vector{d.start, right}, 0}})
	distancesFromEnd := d.dijkstra([]pqItem{
		{vector{d.end, up}, 0},
		{vector{d.end, down}, 0},
		{vector{d.end, right}, 0},
		{vector{d.end, left}, 0},
	})

	target := slices.Min(distancesFromStart[d.end.y][d.end.x])
	viablePositions := make(map[position]bool)

	for y := range len(d.area) {
		for x := range d.area[y] {
			if d.area[y][x] == '#' {
				continue
			}

			if distancesFromStart[y][x][up]+distancesFromEnd[y][x][down] == target ||
				distancesFromStart[y][x][down]+distancesFromEnd[y][x][up] == target ||
				distancesFromStart[y][x][left]+distancesFromEnd[y][x][right] == target ||
				distancesFromStart[y][x][right]+distancesFromEnd[y][x][left] == target {
				viablePositions[position{x, y}] = true
			}

			if distancesFromStart[y][x][up]+distancesFromEnd[y][x][right]+1000 == target ||
				distancesFromStart[y][x][up]+distancesFromEnd[y][x][left]+1000 == target ||
				distancesFromStart[y][x][down]+distancesFromEnd[y][x][right]+1000 == target ||
				distancesFromStart[y][x][down]+distancesFromEnd[y][x][left]+1000 == target ||
				distancesFromStart[y][x][left]+distancesFromEnd[y][x][up]+1000 == target ||
				distancesFromStart[y][x][left]+distancesFromEnd[y][x][down]+1000 == target ||
				distancesFromStart[y][x][right]+distancesFromEnd[y][x][up]+1000 == target ||
				distancesFromStart[y][x][right]+distancesFromEnd[y][x][down]+1000 == target {
				viablePositions[position{x, y}] = true
			}
		}
	}

	count := 0
	for range viablePositions {
		count++
	}

	return fmt.Sprintf("%d\n", count)
}
