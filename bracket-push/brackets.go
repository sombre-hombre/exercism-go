package brackets

var b = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

// Bracket verifies that all the pairs of a brackets, braces and parentheses are matched and nested correctly
// BenchmarkBracket-8   	 1000000	      1352 ns/op	      96 B/op	      12 allocs/op
func Bracket(s string) (bool, error) {
	var st stack
	for _, c := range []byte(s) {
		switch c {
		case '{', '(', '[':
			st.Push(c)
		case '}', ')', ']':
			ok, cb := st.Pop()
			if !ok || c != b[cb] {
				return false, nil
			}
		default:
			continue
		}
	}

	return len(st) == 0, nil
}

type stack []byte

func (s *stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *stack) Pop() (ok bool, c byte) {
	l := len(*s)
	if l > 0 {
		c = (*s)[l-1]
		ok = true
		*s = (*s)[:l-1]
	}
	return
}
