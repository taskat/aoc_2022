package day20

import (
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	data := parseInput(input)
	ring := newRing(data)
	ring.moveAll()
	return fmt.Sprintf("%d", ring.sumOfCoord())
}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	data := make([]int, len(lines))
	var err error
	for i, line := range lines {
		data[i], err = strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
	}
	return data
}

type element struct {
	data int
	index int
	next *element
	prev *element
}

func (e *element) move(length int) {
	amount := e.data % (length - 1)
	if amount == 0 {
		return
	}
	oldPrev := e.prev
	oldNext := e.next
	oldPrev.next = oldNext
	oldNext.prev = oldPrev
	newPrev := e.prev
	newNext := e.next
	if amount > 0 {
		for i := 0; i < amount; i++ {
			newPrev = newPrev.next
			newNext = newNext.next
		}
	} else {
		for i := 0; i > amount; i-- {
			newPrev = newPrev.prev
			newNext = newNext.prev
		}
	}
	newPrev.next = e
	e.prev = newPrev
	newNext.prev = e
	e.next = newNext
}

func (e *element) String() string {
	data := make([]string, 0)
	for iter, i := e, 0; iter != e || i == 0; iter, i = iter.next, 1 {
		data = append(data, fmt.Sprintf("data: %d, index: %d", iter.data, iter.index))
	}
	return strings.Join(data, "\n")
}

type ring struct {
	head *element
	len int
}

func newRing(data []int) *ring {
	var head, prev *element
	for i, d := range data {
		e := element{d, i, nil, prev}
		if prev != nil {
			prev.next = &e
		} else {
			head = &e
		}
		prev = &e
	}
	head.prev = prev
	prev.next = head
	return &ring{head, len(data)}
}

func (r *ring) find(i int) *element {
	for iter, j := r.head, 0; iter != r.head || j == 0; iter, j = iter.next, 1 {
		if iter.index == i {
			return iter
		}
	}
	return nil
}

func (r *ring) moveAll() {
	for i := 0; i < r.len; i++ {
		elem := r.find(i)
		elem.move(r.len)
	}
}

func (r *ring) zero() *element {
	for iter, i := r.head, 0; iter != r.head || i == 0; iter, i = iter.next, 1 {
		if iter.data == 0 {
			return iter
		}
	}
	return nil
}

func (r *ring) sumOfCoord() int {
	zero := r.zero()
	sum := 0
	for iter, i := zero.next, 1; i < 3001; iter, i = iter.next, i + 1 {
		if i % 1_000 == 0 {
			sum += iter.data
		}
	}
	return sum
}

func (r *ring) String() string {
	data := make([]string, 0, r.len)
	for iter, i := r.head, 0; iter != r.head || i == 0; iter, i = iter.next, 1 {
		data = append(data, fmt.Sprint(iter.data))
	}
	return "[" + strings.Join(data, ", ") + "]"
}

const key = 811589153

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	data := parseInput(input)
	decrypt(data, key)
	ring := newRing(data)
	for i := 0; i < 10; i++ {
		ring.moveAll()
	}
	return fmt.Sprintf("%d", ring.sumOfCoord())
}

func decrypt(data []int, key int) {
	for i, d := range data {
		data[i] = d * key
	}
}
