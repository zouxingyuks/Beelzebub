package algorithm

type Stack struct {
	length uint64
	data   string
	next   *Stack
}

// NewStack New 初始化
func NewStack() Stack {
	return Stack{
		length: 0,
		data:   "",
		next:   nil,
	}

}
func (s *Stack) Size() uint64 {
	return s.length
}

// Empty 判断栈是否为空
func (s *Stack) Empty() bool {
	return s.length == 0 && s.next == nil
}

func (s *Stack) Pop() {
	*s = *s.next
}

func (s *Stack) Push(i string) {
	temp := *s
	*s = Stack{
		length: s.length + 1,
		data:   i,
		next:   &temp,
	}
}

func (s *Stack) Top() string {
	if s.Empty() != false {
		return "Stack is empty"
	}
	return s.data
}

func (s *Stack) Clear() {
	for !s.Empty() {
		s.Pop()
	}

}
