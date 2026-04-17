package lc75

func removeStars(s string) string {
	var st Stack

	for _, v := range s {
		if v == '*' {
			st.Pop()
		} else {
			st.Push(v)
		}
	}

	return string(st.items)
}

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
