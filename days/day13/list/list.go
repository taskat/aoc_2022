package list

import (
	"aoc_2022/days/day13/value"
	"strings"
)

type List []value.Value

func NewList(values ...value.Value) List {
	return List(values)
}

func (l List) Compare(other value.Value) value.Order {
	otherList, ok := other.(List)
	if !ok {
		return l.Compare(other.ToList())
	}
	for i := 0; i < len(l) && i < len(otherList); i++ {
		compare := l[i].Compare(otherList[i])
		if compare != value.DontKnow {
			return compare
		}
	}
	if len(l) < len(otherList) {
		return value.Right
	}
	if len(l) > len(otherList) {
		return value.Wrong
	}
	return value.DontKnow
}

func (l List) String() string {
	parts := make([]string, len(l))
	for i, v := range l {
		parts[i] = v.String()
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func (l List) ToList() value.Value {
	return l
}