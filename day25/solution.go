package day25

import "fmt"

const HEIGHT int = 7
const WIDTH int = 5

type Day25 struct {
	pins [][]int
	keys [][]int
}

func makePins(pattern []string) (pins []int) {
	for i := range WIDTH {
		c := 0
		for j := range HEIGHT {
			if pattern[j][i] == '.' {
				break
			}
			c++
		}
		pins = append(pins, c)
	}

	return pins
}

func makeKey(pattern []string) (key []int) {
	for i := range WIDTH {
		c := 0
		for j := range HEIGHT {
			if pattern[j][i] == '#' {
				break
			}
			c++
		}
		key = append(key, HEIGHT-c)
	}

	return key
}

func (d *Day25) parseIntoPinsAndKeys(pattern []string) {
	if pattern[0][0] == '#' {
		d.pins = append(d.pins, makePins(pattern))
	} else {
		d.keys = append(d.keys, makeKey(pattern))
	}
}

func (d *Day25) Init(lines []string) {
	pattern := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			d.parseIntoPinsAndKeys(pattern)
			pattern = make([]string, 0)
		} else {
			pattern = append(pattern, line)
		}
	}
	d.parseIntoPinsAndKeys(pattern)
}

func (d *Day25) SolveSimple() string {
	compatible := 0
	for _, pin := range d.pins {
		for _, key := range d.keys {
			compatible++
			for i := range WIDTH {
				if pin[i]+key[i] > HEIGHT {
					compatible--
					break
				}
			}
		}
	}
	return fmt.Sprint(compatible)
}

func (d *Day25) SolveAdvanced() string {
	return ""
}
