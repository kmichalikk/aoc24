package day9

import (
	"fmt"
)

type fragment struct {
	id            int
	start, length int
}

type Day9 struct {
	fragments         [][]fragment
	fragmentsAdvanced [][]fragment
	free              []fragment
	freeAdvanced      []fragment
}

func (d *Day9) Init(lines []string) {
	isFree := false
	offset := 0
	d.fragments = make([][]fragment, 0)
	d.fragmentsAdvanced = make([][]fragment, 0)
	d.free = make([]fragment, 0)
	d.freeAdvanced = make([]fragment, 0)
	for _, ch := range lines[0] {
		s := ch - '0'
		if isFree {
			d.free = append(d.free, fragment{0, offset, int(s)})
			d.freeAdvanced = append(d.freeAdvanced, fragment{0, offset, int(s)})
		} else {
			d.fragments = append(d.fragments, make([]fragment, 1))
			d.fragments[len(d.fragments)-1][0] = fragment{len(d.fragments) - 1, offset, int(s)}
			d.fragmentsAdvanced = append(d.fragmentsAdvanced, make([]fragment, 1))
			d.fragmentsAdvanced[len(d.fragments)-1][0] = fragment{len(d.fragmentsAdvanced) - 1, offset, int(s)}
		}
		offset += int(s)
		isFree = !isFree
	}
}

func (d *Day9) SolveSimple() string {
	i := 0
	start, space := 0, 0
	for j := len(d.fragments) - 1; j > 0; j-- {
		frag := d.fragments[j][0]
		d.fragments[j] = make([]fragment, 0)
		for frag.length > 0 {
			if space == 0 {
				if i == len(d.free) {
					break
				}
				start = d.free[i].start
				space = d.free[i].length
				i++
			}
			alloc := min(frag.length, space)
			d.fragments[j] = append(d.fragments[j], fragment{frag.id, start, alloc})
			space -= alloc
			start += alloc
			frag.length -= alloc
		}
		if i == len(d.free) {
			if frag.length > 0 {
				d.fragments[j] = append(d.fragments[j], frag)
			}
			break
		}
		d.free = d.free[:len(d.free)-1]
	}

	checksum := 0
	for _, fragments := range d.fragments {
		for _, frag := range fragments {
			checksum += frag.id * frag.length * (2*frag.start + frag.length - 1) / 2
		}
	}

	return fmt.Sprintf("%d", checksum)
}

func (d *Day9) SolveAdvanced() string {
	for i := len(d.fragmentsAdvanced) - 1; i > 0; i-- {
		for j := 0; j < len(d.freeAdvanced) && d.freeAdvanced[j].start < d.fragmentsAdvanced[i][0].start; j++ {
			if d.freeAdvanced[j].length >= d.fragmentsAdvanced[i][0].length {
				d.fragmentsAdvanced[i][0].start = d.freeAdvanced[j].start
				d.freeAdvanced[j].length -= d.fragmentsAdvanced[i][0].length
				d.freeAdvanced[j].start += d.fragmentsAdvanced[i][0].length
			}
		}
	}

	checksum := 0
	for _, fragments := range d.fragmentsAdvanced {
		for _, frag := range fragments {
			checksum += frag.id * frag.length * (2*frag.start + frag.length - 1) / 2
		}
	}

	return fmt.Sprintf("%d", checksum)
}
