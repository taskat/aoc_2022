package config

import "strconv"

type input struct {
	real bool
	test int
}

func newRealData() *input {
	return &input{real: true}
}

func newTestData(test int) *input {
	return &input{test: test}
}

func (i *input) String() string {
	if i.real {
		return ""
	} else {
		return strconv.Itoa(i.test)
	}
}