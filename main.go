package main

import (
	"aoc24/day15"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day15/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	problem := day15.Day15{}
	problem.Init(lines)

	fmt.Println(problem.SolveSimple())
	fmt.Println(problem.SolveAdvanced())
}
