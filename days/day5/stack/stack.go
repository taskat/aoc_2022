package stack

type Stack struct {
	data []rune
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *Stack) PushAll(r []rune) {
	s.data = append(s.data, r...)
}

func (s *Stack) pop() rune {
	if len(s.data) == 0 {
		panic("Stack is empty")
	}
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *Stack) Pop(n int) []rune {
	data := make([]rune, n)
	for i := 0; i < n; i++ {
		data[i] = s.pop()
	}
	return data
}

func (s *Stack) Top() rune {
	if len(s.data) == 0 {
		panic("Stack is empty")
	}
	return s.data[len(s.data)-1]
}