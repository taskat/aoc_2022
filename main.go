package main

import (
	"aoc_2022/config"
	"aoc_2022/days/day1"
	"aoc_2022/solver"
	"fmt"
	"os"
)

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
	input := cfg.GetInputData()
	fmt.Printf("Start solving day %d, part %d with %s input...\n", cfg.GetDay(), cfg.GetPart(), cfg.GetInputType())
	solution := solve(cfg, input)
	fmt.Printf("Solved! Solution is: %s\n", solution)
}