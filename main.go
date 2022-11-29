package main

import (
	"aoc_2022/config"
	"aoc_2022/days/day1"
	"aoc_2022/solver"
	"fmt"
	"os"
)

func getInput(cfg *config.Config) string {
	fileName := cfg.GetInputFilename()
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	return string(data)
}

func getSolver(cfg *config.Config) solver.Solver {
	switch cfg.GetDay() {
	case 1:
		return &day1.Solver{}
	default:
		panic("Day not implemented yet")
	}
}

func solve(cfg *config.Config, input string) string {
	solver := getSolver(cfg)
	if cfg.GetPart() == 1 {
		return solver.SolvePart1(input, cfg.GetExtraParams()...)
	} else {
		return solver.SolvePart2(input, cfg.GetExtraParams()...)
	}
}

func main() {
	cfg := config.ParseConfig()
	if cfg == nil {
		os.Exit(1)
	}
	input := getInput(cfg)
	fmt.Printf("Start solving day %d, part %d with %s input...\n", cfg.GetDay(), cfg.GetPart(), cfg.GetInput())
	solution := solve(cfg, input)
	fmt.Printf("Solved! Solution is: %s\n", solution)
}