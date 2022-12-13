package integer

import (
	"aoc_2022/days/day13/list"
	"aoc_2022/days/day13/value"
	"fmt"
)

type Integer int

func NewInteger(value int) Integer {
	return Integer(value)
}

func (i Integer) Compare(other value.Value) value.Order {
	otherInteger, ok := other.(Integer)
	if ok {
		if i == otherInteger {
			return value.DontKnow
		} else if i < otherInteger {
			return value.Right
		} else {
			return value.Wrong
		}
	}
	return i.ToList().Compare(other)
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (i Integer) ToList() value.Value {
	return list.NewList(i)
}