package main

type MinStack struct {
	stack []int
	mix   []int
	deep  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mix:   []int{},
		deep:  0,
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	s.mix = append(s.mix, x)
	s.deep++

	for i := s.deep - 1; i > 0; i-- {
		if s.mix[i] > s.mix[i-1] {
			s.mix[i], s.mix[i-1] = s.mix[i-1], s.mix[i]
		} else {
			break
		}
	}

}

func (s *MinStack) Pop() {

	if s.deep == 0 {
		return
	}

	p := s.stack[s.deep-1]
	s.stack = s.stack[:s.deep-1]
	for i := 0; i <= s.deep; i++ {
		if s.mix[i] == p {
			s.mix = append(s.mix[:i], s.mix[i+1:]...)
			break
		}
	}
	s.deep--
}

func (s *MinStack) Top() int {
	return s.stack[s.deep-1]
}

func (s *MinStack) GetMin() int {
	return s.mix[s.deep-1]
}

func main() {

}
