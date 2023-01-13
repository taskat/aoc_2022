package main

import (
	"aoc_2022/config"
	"aoc_2022/days/day1"
	"aoc_2022/days/day2"
	"aoc_2022/days/day3"
	"aoc_2022/days/day4"
	"aoc_2022/days/day5"
	"aoc_2022/days/day6"
	"aoc_2022/days/day7"
	"aoc_2022/days/day8"
	"aoc_2022/days/day9"
	"aoc_2022/days/day10"
	"aoc_2022/days/day11"
	"aoc_2022/days/day12"
	"aoc_2022/days/day13"
	"aoc_2022/days/day14"
	"aoc_2022/days/day15"
	"aoc_2022/days/day16"
	"aoc_2022/days/day17"
	"aoc_2022/days/day18"
	"aoc_2022/days/day19"
	"aoc_2022/days/day20"
	"aoc_2022/solver"
	"fmt"
	"os"
)

type Config interface {
	GetDay() int
	GetPart() int
	GetInputType() string
	GetInputData() string
	GetExtraParams() []interface{}
}

func getSolver(cfg Config) solver.Solver {
	switch cfg.GetDay() {
	case 1:
		return &day1.Solver{}
	case 2:
		return &day2.Solver{}
	case 3:
		return &day3.Solver{}
	case 4:
		return &day4.Solver{}
	case 5:
		return &day5.Solver{}
	case 6:
		return &day6.Solver{}
	case 7:
		return &day7.Solver{}
	case 8:
		return &day8.Solver{}
	case 9:
		return &day9.Solver{}
	case 10:
		return &day10.Solver{}
	case 11:
		return &day11.Solver{}
	case 12:
		return &day12.Solver{}
	case 13:
		return &day13.Solver{}
	case 14:
		return &day14.Solver{}
	case 15:
		return &day15.Solver{}
	case 16:
		return &day16.Solver{}
	case 17:
		return &day17.Solver{}
	case 18:
		return &day18.Solver{}
	case 19:
		return &day19.Solver{}
	case 20:
		return &day20.Solver{}
	default:
		panic("Day not implemented yet")
	}
}

func solve(cfg Config, input string) string {
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