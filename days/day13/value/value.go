package value

import "fmt"

type Order int

const (
	Right Order = iota
	Wrong
	DontKnow
)

func (o Order) String() string {
	switch o {
	case Right:
		return "Right"
	case Wrong:
		return "Wrong"
	case DontKnow:
		return "DontKnow"
	default:
		return "?"
	}
}

type Value interface {
	Compare(other Value) Order
	ToList() Value
	fmt.Stringer
}