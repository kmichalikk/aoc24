package common

type Problem interface {
	Init(lines []string)
	SolveSimple() string
	SolveAdvanced() string
}
