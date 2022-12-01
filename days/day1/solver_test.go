package day1

import (
	"aoc_2022/config"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day = 1

func TestSolvePart1(t *testing.T) {
	testCases := []struct {
		name string
		input config.Input
		extraParams []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "24000"},
		{"Real", *config.NewRealInput(), nil, "75622"},
	}
	fmt.Println(os.Args)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart1(cfg.GetInputData(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}

func TestSolvePart2(t *testing.T) {
	testCases := []struct {
		name string
		input config.Input
		extraParams []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "45000"},
		{"Real", *config.NewRealInput(), nil, "213159"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			solver := &Solver{}
			cfg := config.NewConfigForTest(config.NewConfig(day, 0, tc.input))
			solution := solver.SolvePart2(cfg.GetInputData(), tc.extraParams...)
			assert.Equal(t, tc.expectedValue, solution)
		})
	}
}