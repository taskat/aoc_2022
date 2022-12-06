package day6

import (
	"aoc_2022/config"
	"testing"
	"github.com/stretchr/testify/assert"
)

var day = 6

func TestSolvePart1(t *testing.T) {
	testCases := []struct {
		name string
		input config.Input
		extraParams []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "7"},
		{"Test 2", *config.NewTestInput(2), nil, "5"},
		{"Test 3", *config.NewTestInput(3), nil, "6"},
		{"Test 4", *config.NewTestInput(4), nil, "10"},
		{"Test 5", *config.NewTestInput(5), nil, "11"},
		{"Real", *config.NewRealInput(), nil, "1343"},
	}
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
		{"Test 1", *config.NewTestInput(1), nil, "19"},
		{"Test 2", *config.NewTestInput(2), nil, "23"},
		{"Test 3", *config.NewTestInput(3), nil, "23"},
		{"Test 4", *config.NewTestInput(4), nil, "29"},
		{"Test 5", *config.NewTestInput(5), nil, "26"},
		{"Real", *config.NewRealInput(), nil, "2193"},
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