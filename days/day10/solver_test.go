package day10

import (
	"aoc_2022/config"
	"testing"
	"github.com/stretchr/testify/assert"
)

var day = 10

func TestSolvePart1(t *testing.T) {
	testCases := []struct {
		name string
		input config.Input
		extraParams []any
		expectedValue string
	}{
		{"Test 1", *config.NewTestInput(1), nil, "13140"},
		{"Real", *config.NewRealInput(), nil, "13820"},
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
