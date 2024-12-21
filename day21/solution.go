package day21

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type action struct {
	code   int32
	target string
}

type Day21 struct {
	codes    []string
	graph    map[string][]action
	shortest map[int32]map[int32][]string
}

func (d *Day21) Init(lines []string) {
	d.codes = lines
	d.graph = map[string][]action{
		"AA": {{'A', "AA"}, {'v', "A>"}, {'<', "A^"}},
		"A^": {{'>', "AA"}, {'v', "Av"}},
		"Av": {{'^', "A^"}, {'<', "A<"}, {'>', "A>"}, {'A', ">v"}},
		"A<": {{'>', "A>"}, {'A', "^<"}},
		"A>": {{'<', "Av"}, {'^', "AA"}},
		"^A": {{'A', "^A"}, {'v', "^>"}, {'<', "^^"}},
		"^^": {{'>', "^A"}, {'v', "^v"}},
		"^v": {{'A', "vv"}, {'<', "^<"}, {'>', "^>"}, {'^', "^^"}},
		"^<": {{'>', "^v"}},
		"^>": {{'A', "A>"}, {'<', "^v"}, {'^', "^A"}},
		"vA": {{'A', "vA"}, {'v', "v>"}, {'<', "v^"}},
		"v^": {{'A', "^^"}, {'v', "vv"}, {'>', "vA"}},
		"vv": {{'^', "v^"}, {'<', "v<"}, {'>', "v>"}},
		"v<": {{'A', "<<"}, {'>', "vv"}},
		"v>": {{'A', ">>"}, {'<', "vv"}, {'^', "vA"}},
		"<A": {{'A', "<A"}, {'<', "<^"}, {'v', "<>"}},
		"<^": {{'>', "<A"}, {'v', "<v"}},
		"<v": {{'^', "<^"}, {'<', "<<"}, {'>', "<>"}},
		"<<": {{'>', "<v"}},
		"<>": {{'<', "<v"}, {'A', "v>"}, {'^', "<A"}},
		">A": {{'A', ">A"}, {'<', ">^"}, {'v', ">>"}},
		">^": {{'A', "A^"}, {'>', ">A"}, {'v', ">v"}},
		">v": {{'^', ">^"}, {'<', "><"}, {'>', ">>"}},
		"><": {{'A', "v<"}, {'>', ">v"}},
		">>": {{'<', ">v"}, {'^', ">A"}},
	}

	d.shortest = map[int32]map[int32][]string{
		'A': {
			'A': {},
			'0': {"<A"},
			'1': {"^A", "<A", "<A"},
			'2': {"<A", "^A"},
			'3': {"^A"},
			'4': {"^A", "^A", "<A", "<A"},
			'5': {"<A", "^A", "^A"},
			'6': {"^A", "^A"},
			'7': {"^A", "^A", "^A", "<A", "<A"},
			'8': {"<A", "^A", "^A", "^A"},
			'9': {"^A", "^A", "^A"},
		},
		'0': {
			'A': {">A"},
			'0': {},
			'1': {"^A", "<A"},
			'2': {"^A"},
			'3': {"^A", ">A"},
			'4': {"^A", "^A", "<A"},
			'5': {"^A", "^A"},
			'6': {"^A", "^A", ">A"},
			'7': {"^A", "^A", "^A", "<A"},
			'8': {"^A", "^A", "^A"},
			'9': {"^A", "^A", "^A", ">A"},
		},
		'1': {
			'A': {">A", ">A", "vA"},
			'0': {">A", "vA"},
			'1': {},
			'2': {">A"},
			'3': {">A", ">A"},
			'4': {"^A"},
			'5': {"^A", ">A"},
			'6': {"^A", ">A", ">A"},
			'7': {"^A", "^A"},
			'8': {"^A", "^A", ">A"},
			'9': {"^A", "^A", ">A", ">A"},
		},
		'2': {
			'A': {">A", "vA"},
			'0': {"vA"},
			'1': {"<A"},
			'2': {},
			'3': {">A"},
			'4': {"<A", "^A"},
			'5': {"^A"},
			'6': {"^A", ">A"},
			'7': {"<A", "^A", "^A"},
			'8': {"^A", "^A"},
			'9': {"^A", "^A", ">A"},
		},
		'3': {
			'A': {"vA"},
			'0': {"vA", "<A"},
			'1': {"<A", "<A"},
			'2': {"<A"},
			'3': {},
			'4': {"<A", "<A", "^A"},
			'5': {"<A", "^A"},
			'6': {"^A"},
			'7': {"<A", "<A", "^A", "^A"},
			'8': {"^A", "^A", "<A"},
			'9': {"^A", "^A"},
		},
		'4': {
			'A': {">A", ">A", "vA", "vA"},
			'0': {">A", "vA", "vA"},
			'1': {"vA"},
			'2': {">A", "vA"},
			'3': {">A", ">A", "vA"},
			'4': {},
			'5': {">A"},
			'6': {">A", ">A"},
			'7': {"^A"},
			'8': {"^A", ">A"},
			'9': {"^A", ">A", ">A"},
		},
		'5': {
			'A': {">A", "vA", "vA"},
			'0': {"vA", "vA"},
			'1': {"vA", "<A"},
			'2': {"vA"},
			'3': {"vA", ">A"},
			'4': {"<A"},
			'5': {},
			'6': {">A"},
			'7': {"<A", "^A"},
			'8': {"^A"},
			'9': {"^A", ">A"},
		},
		'6': {
			'A': {"vA", "vA"},
			'0': {"vA", "vA", "<A"},
			'1': {"vA", "<A", "<A"},
			'2': {"vA", "<A"},
			'3': {"vA"},
			'4': {"<A", "<A"},
			'5': {"<A"},
			'6': {},
			'7': {"<A", "<A", "^A"},
			'8': {"<A", "^A"},
			'9': {"^A"},
		},
		'7': {
			'A': {">A", ">A", "vA", "vA", "vA"},
			'0': {">A", "vA", "vA", "vA"},
			'1': {"vA", "vA"},
			'2': {">A", "vA", "vA"},
			'3': {">A", ">A", "vA", "vA"},
			'4': {"vA"},
			'5': {">A", "vA"},
			'6': {">A", ">A", "vA"},
			'7': {},
			'8': {">A"},
			'9': {">A", ">A"},
		},
		'8': {
			'A': {">A", "vA", "vA", "vA"},
			'0': {"vA", "vA", "vA"},
			'1': {"vA", "vA", "<A"},
			'2': {"vA", "vA"},
			'3': {">A", "vA", "vA"},
			'4': {"vA", "<A"},
			'5': {"vA"},
			'6': {">A", "vA"},
			'7': {"<A"},
			'8': {},
			'9': {">A"},
		},
		'9': {
			'A': {"vA", "vA", "vA"},
			'0': {"vA", "vA", "vA", "<a"},
			'1': {"vA", "vA", "<A", "<A"},
			'2': {"vA", "vA", "<A"},
			'3': {"vA", "vA"},
			'4': {"vA", "<A", "<A"},
			'5': {"vA", "<A"},
			'6': {"vA"},
			'7': {"<A", "<A"},
			'8': {"<A"},
			'9': {},
		},
	}
}

type item struct {
	w int
	v string
	p string
}

func (d *Day21) bfs(start, end string) int {
	visited := make(map[string]bool)
	parent := make(map[string]string)
	queue := []item{{0, start, "@"}}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		if visited[u.v] {
			continue
		}
		visited[u.v] = true
		parent[u.v] = u.p

		if u.v == end {
			break
		}

		for _, a := range d.graph[u.v] {
			if visited[a.target] {
				continue
			}

			queue = append(queue, item{u.w + 1, a.target, u.v})
		}
	}

	b := strings.Builder{}
	b.WriteString("A")
	s := end
	ps := parent[end]

	if ps == "" {
		log.Panicln("path not found")
		return -1
	}

	for ps != "@" {
		for _, c := range d.graph[ps] {
			if c.target == s {
				b.WriteString(string(c.code))
				break
			}
		}

		s = ps
		ps = parent[s]
	}

	str := []byte(b.String())
	slices.Reverse(str)

	fmt.Print(string(str))
	return len(str)
}

func (d *Day21) SolveSimple() string {
	total := 0
	for _, code := range d.codes {
		state := "AA"
		lastChar := 'A'
		length := 0
		for _, ch := range code {
			state = "AA"
			for _, p := range d.shortest[lastChar][ch] {
				length += d.bfs(state, p)
				state = p
			}
			length += d.bfs(state, "AA")
			lastChar = ch
		}
		fmt.Println("\n", code, length)
		val, _ := strconv.ParseInt(code[:3], 10, 32)
		total += length * int(val)
	}

	return fmt.Sprint(total)
}

func (d *Day21) SolveAdvanced() string {
	return ""
}

/*
029A: <vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A
980A: <v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A
179A: <v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
456A: <v<A>>^AA<vA<A>>^AAvAA<^A>A<vA>^A<A>A<vA>^A<A>A<v<A>A>^AAvA<^A>A
379A: <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A

*/
