package day13

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Day13 struct {
	arcades []arcade
}

type vector struct {
	x, y int64
}

func (v vector) eq(other vector) bool {
	return v.x == other.x && v.y == other.y
}

func (v vector) add(other vector) vector {
	return vector{v.x + other.x, v.y + other.y}
}

type arcade struct {
	a, b  vector
	prize vector
}

func (d *Day13) Init(lines []string) {
	buttonRegex, err := regexp.Compile("Button [AB]: X\\+(\\d+), Y\\+(\\d+)")
	if err != nil {
		log.Fatal("bad button regex")
	}
	prizeRegex, err := regexp.Compile("Prize: X=(\\d+), Y=(\\d+)")
	if err != nil {
		log.Fatal("bad prize regex")
	}
	d.arcades = make([]arcade, 0)
	for i := 0; i < len(lines); i += 4 {
		aMatch := buttonRegex.FindStringSubmatch(lines[i])
		bMatch := buttonRegex.FindStringSubmatch(lines[i+1])
		prizeMatch := prizeRegex.FindStringSubmatch(lines[i+2])
		ax, _ := strconv.Atoi(aMatch[1])
		ay, _ := strconv.Atoi(aMatch[2])
		bx, _ := strconv.Atoi(bMatch[1])
		by, _ := strconv.Atoi(bMatch[2])
		prizex, _ := strconv.Atoi(prizeMatch[1])
		prizey, _ := strconv.Atoi(prizeMatch[2])
		d.arcades = append(
			d.arcades,
			arcade{
				vector{int64(ax), int64(ay)},
				vector{int64(bx), int64(by)},
				vector{int64(prizex), int64(prizey)},
			},
		)
	}
}

func (d *Day13) SolveSimple() string {
	var search func(position vector, memo *map[vector]int, i, aDepth, bDepth int) int
	search = func(position vector, memo *map[vector]int, i, aDepth, bDepth int) int {
		switch {
		case (*memo)[position] != 0:
			return (*memo)[position]
		case position.eq(d.arcades[i].prize):
			(*memo)[position] = 0
			return (*memo)[position]
		case aDepth == 100 || bDepth == 100:
			(*memo)[position] = 999
			return (*memo)[position]
		}

		(*memo)[position] = min(
			search(position.add(d.arcades[i].a), memo, i, aDepth+1, bDepth)+3,
			search(position.add(d.arcades[i].b), memo, i, aDepth, bDepth+1)+1,
		)
		return (*memo)[position]
	}

	total := 0
	for i := range d.arcades {
		memo := make(map[vector]int)
		tokens := search(vector{0, 0}, &memo, i, 0, 0)
		if tokens < 999 {
			total += tokens
		}
	}

	return fmt.Sprintf("%d\n", total)
}

func (d *Day13) SolveAdvanced() string {
	return ""
}
