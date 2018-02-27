package secret

var code = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake converts a number to a secret handshake
// BenchmarkHandshake-8   	  200000	      7940 ns/op	    3216 B/op	      81 allocs/op
func HandshakeV2(n uint) []string {
	var result = make([]string, 0, 4)
	reverse := n&16 > 0
	for i, word := range code {
		if (n & (1 << uint(i))) != 0 {
			if reverse {
				result = append([]string{word}, result...)
			} else {
				result = append(result, word)
			}
		}
	}

	return result
}

// Handshake converts a number to a secret handshake
// BenchmarkHandshake-8   	  300000	      4199 ns/op	    2048 B/op	      32 allocs/op
func Handshake(n uint) []string {
	var result = make([]string, 0, 4)
	if n&16 > 0 {
		for i := len(code) - 1; i >= 0; i-- {
			if (n & (1 << uint(i))) != 0 {
				result = append(result, code[i])
			}
		}
	} else {
		for i := 0; i < len(code); i++ {
			if (n & (1 << uint(i))) != 0 {
				result = append(result, code[i])
			}
		}
	}

	return result
}
