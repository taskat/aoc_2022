package day3

import (
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	rucksacks := parseRucksacks(input)
	commons := getCommons(rucksacks)
	return strconv.Itoa(sumPriorities(commons))
}

func parseRucksacks(input string) []rucksack {
	lines := strings.Split(input, "\n")
	rucksacks := make([]rucksack, len(lines))
	for i, line := range lines {
		rucksacks[i] = newRucksack(line)
	}
	return rucksacks
}

type rucksack [2]string

func newRucksack(items string) rucksack {
	compartment1 := items[:len(items)/2]
	compartment2 := items[len(items)/2:]
	return rucksack([2]string{compartment1, compartment2})
}

func (r rucksack) getCommonItem() rune {
	for _, c1 := range r[0] {
		for _, c2 := range r[1] {
			if c1 == c2 {
				return c1
			}
		}
	}
	panic("No common item found")
}

func (r rucksack) items() string {
	return r[0] + r[1]
}

func getCommons(rucksacks []rucksack) []rune {
	commonItems := make([]rune, len(rucksacks))
	for i, r := range rucksacks {
		commonItems[i] = r.getCommonItem()
	}
	return commonItems
}

func sumPriorities(commons []rune) int {
	sum := 0
	for _, c := range commons {
		sum += getPriority(c)
	}
	return sum
}

func getPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a') + 1
	}
	return int(item - 'A') + 27
}


func (*Solver) SolvePart2(input string, extraParams ...any) string {
	rucksacks := parseRucksacks(input)
	groups := createGroups(rucksacks)
	badges := getBadges(groups)
	return strconv.Itoa(sumPriorities(badges))
}

func createGroups(rucksacks []rucksack) []group {
	groups := make([]group, len(rucksacks)/3)
	for i := 0; i < len(rucksacks); i += 3 {
		groups[i/3] = group{rucksacks[i], rucksacks[i+1], rucksacks[i+2]}
	}
	return groups
}

type group [3]rucksack

func (g group) getBadge() rune {
	commons := make([]rune, 0)
	for _, c1 := range g[0].items() {
		for _, c2 := range g[1].items() {
			if c1 == c2 {
				commons = append(commons, c1)
			}
		}
	}
	for _, c1 := range g[2].items() {
		for _, c2 := range commons {
			if c1 == c2 {
				return c1
			}
		}
	}
	panic("No common item found")
}

func getBadges(groups []group) []rune {
	badges := make([]rune, len(groups))
	for i, g := range groups {
		badges[i] = g.getBadge()
	}
	return badges
}