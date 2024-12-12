package main

import (
	"aoc24/day12"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day12/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	problem := day12.Day12{}
	problem.Init(lines)

	fmt.Println(problem.SolveSimple())
	fmt.Println(problem.SolveAdvanced())
}
