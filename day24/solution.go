package day24

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	OR = iota
	AND
	XOR
	LITERAL
)

type expr struct {
	op   int
	a, b string
	val  int
}

type Day24 struct {
	exprs map[string]expr
}

func (d *Day24) Init(lines []string) {
	literalsRegexp := regexp.MustCompile("(\\w{3}): ([01])")
	gatesRegexp := regexp.MustCompile("(\\w{3}) (XOR|LITERAL|AND|OR) (\\w{3}) -> (\\w{3})")
	d.exprs = make(map[string]expr)
	i := 0
	for lines[i] != "" {
		matches := literalsRegexp.FindStringSubmatch(lines[i])
		val := 0
		if matches[2] == "1" {
			val = 1
		}
		d.exprs[matches[1]] = expr{LITERAL, "", "", val}
		i++
	}
	i++
	for i < len(lines) {
		matches := gatesRegexp.FindStringSubmatch(lines[i])
		switch matches[2] {
		case "AND":
			d.exprs[matches[4]] = expr{AND, matches[1], matches[3], -1}
		case "OR":
			d.exprs[matches[4]] = expr{OR, matches[1], matches[3], -1}
		case "XOR":
			d.exprs[matches[4]] = expr{XOR, matches[1], matches[3], -1}
		}
		i++
	}
}

func (d *Day24) Eval(e expr) int {
	if e.val != -1 {
		return e.val
	}

	switch e.op {
	case OR:
		e.val = d.Eval(d.exprs[e.a]) | d.Eval(d.exprs[e.b])
	case AND:
		e.val = d.Eval(d.exprs[e.a]) & d.Eval(d.exprs[e.b])
	case XOR:
		e.val = d.Eval(d.exprs[e.a]) ^ d.Eval(d.exprs[e.b])
	}

	return e.val
}

func (d *Day24) SolveSimple() string {
	outputs := make([]string, 0)
	for k := range d.exprs {
		if k[0] == 'z' {
			outputs = append(outputs, k)
		}
	}
	slices.Sort(outputs)
	slices.Reverse(outputs)
	bitsBuilder := strings.Builder{}
	for _, z := range outputs {
		bitsBuilder.WriteString(strconv.Itoa(d.Eval(d.exprs[z])))
	}

	value, _ := strconv.ParseInt(bitsBuilder.String(), 2, 64)

	return fmt.Sprint(value)
}

func (d *Day24) SolveAdvanced() string {
	return ""
}
