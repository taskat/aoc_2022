package day22

import (
	"aoc_2022/config"
	"testing"
	"github.com/stretchr/testify/assert"
)

var day = 22

func TestSolvePart1(t *testing.T) {
	testCases := []struct {
		name string
		input config.Input
		extraParams []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), []any{"4"}, "6032"},
		{"Real", *config.NewRealInput(), nil, "149250"},
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
		{"Test 1", *config.NewTestInput(1), []any{"4"}, "5031"},
		{"Real", *config.NewRealInput(), nil, "12462"},
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