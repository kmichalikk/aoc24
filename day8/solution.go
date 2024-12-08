package day8

import "fmt"

type vector struct {
	x, y int
}

func (p1 vector) diff(p2 vector) vector {
	return vector{p2.x - p1.x, p2.y - p1.y}
}

func (p1 vector) add(p2 vector) vector {
	return vector{p1.x + p2.x, p1.y + p2.y}
}

func (p1 vector) sub(p2 vector) vector {
	return vector{p1.x - p2.x, p1.y - p2.y}
}

type Day8 struct {
	lines         []string
	antennas      map[int32][]vector
	width, height int
}

func (d *Day8) Init(lines []string) {
	d.antennas = make(map[int32][]vector)
	d.width = len(lines[0])
	d.height = len(lines)
	for i, line := range lines {
		for j, freq := range line {
			if freq >= '0' && freq <= '9' || freq >= 'A' && freq <= 'Z' || freq >= 'a' && freq <= 'z' {
				_, ok := d.antennas[freq]
				if !ok {
					d.antennas[freq] = make([]vector, 0)
				}

				d.antennas[freq] = append(d.antennas[freq], vector{j, i})
			}
		}
	}
}

func (d *Day8) SolveSimple() string {
	antinodes := make(map[vector]bool)
	for _, positions := range d.antennas {
		i := 0
		for i < len(positions)-1 {
			for j := i + 1; j < len(positions); j++ {
				diff := positions[i].diff(positions[j])
				p1 := positions[i].sub(diff)
				p2 := positions[j].add(diff)
				if p1.x >= 0 && p1.x < d.width && p1.y >= 0 && p1.y < d.height {
					antinodes[p1] = true
				}
				if p2.x >= 0 && p2.x < d.width && p2.y >= 0 && p2.y < d.height {
					antinodes[p2] = true
				}
			}
			i++
		}
	}

	return fmt.Sprintf("%d\n", len(antinodes))
}

func (d *Day8) SolveAdvanced() string {
	antinodes := make(map[vector]bool)
	for _, positions := range d.antennas {
		i := 0
		for i < len(positions)-1 {
			for j := i + 1; j < len(positions); j++ {
				diff := positions[i].diff(positions[j])
				antinodes[positions[i]] = true
				antinodes[positions[j]] = true
				pos := positions[i].sub(diff)
				for pos.x >= 0 && pos.x < d.width && pos.y >= 0 && pos.y < d.height {
					antinodes[pos] = true
					pos = pos.sub(diff)
				}
				pos = positions[j].add(diff)
				for pos.x >= 0 && pos.x < d.width && pos.y >= 0 && pos.y < d.height {
					antinodes[pos] = true
					pos = pos.add(diff)
				}
			}
			i++
		}
	}

	return fmt.Sprintf("%d\n", len(antinodes))
}
