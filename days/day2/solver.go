package day2

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	rounds := createRounds(input)
	score := scoreRounds(rounds)
	return strconv.Itoa(score)
}

type rps int

const (
	ROCK rps = iota
	PAPER
	SCISSORS
)

func parseRps(input string) rps {
	switch input {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	}
	panic("Invalid input for parse")
}

func (r rps) score() int {
	switch r {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	}	
	panic("Invalid input for score")
}

type round struct {
	elf rps
	player rps
}

func newRound(elf, player rps) round {
	return round{elf, player}
}

func (r round) isDraw() bool {
	return r.elf == r.player
}

func (r round) isWin() bool {
	return r.player == r.elf + 1 || r.player == r.elf - 2
}

func (r round) score() int {
	score := r.player.score()
	if r.isWin() {
		score += 6
	} else if r.isDraw() {
		score += 3
	}
	return score
}

func createRounds(input string) []round {
	lines := strings.Split(input, "\n")
	rounds := make([]round, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		elf := parseRps(parts[0])
		player := parseRps(parts[1])
		rounds[i] = newRound(elf, player)
	}
	return rounds
}

func scoreRounds(rounds []round) int {
	score := 0
	for _, round := range rounds {
		score += round.score()
	}
	return score
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	rounds := createRoundsByOutcome(input)
	score := scoreRounds(rounds)
	return strconv.Itoa(score)
}

func createRoundsByOutcome(input string) []round {
	lines := strings.Split(input, "\n")
	rounds := make([]round, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		elf := parseRps(parts[0])
		player := choose(elf, parts[1])
		rounds[i] = newRound(elf, player)
	}
	return rounds
}

func choose(elf rps, outcome string) rps {
	switch outcome {
	case "X":
		return (elf + 2) % 3
	case "Y":
		return elf
	case "Z":
		return (elf + 1) % 3
	}
	panic("Invalid input for choose")
}
