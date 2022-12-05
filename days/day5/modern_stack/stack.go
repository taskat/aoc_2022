package modern_stack

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

func (s *Stack) Pop(n int) []rune {
	data := s.data[len(s.data)-n:]
	s.data = s.data[:len(s.data)-n]
	return data
}

func (s *Stack) Top() rune {
	if len(s.data) == 0 {
		panic("Stack is empty")
	}
	return s.data[len(s.data)-1]
}