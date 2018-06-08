package railfence

// Encode implements encoding for the rail fence cipher.
// BenchmarkEncode-8   	 2000000	       632 ns/op	     592 B/op	       9 allocs/op
func Encode(s string, n int) string {
	result := make([]byte, len(s))

	for i, pos := range getKey(n, len(s)) {
		result[i] = s[pos]
	}

	return string(result)
}

// Decode implements decoding for the rail fence cipher.
// BenchmarkDecode-8   	 2000000	       973 ns/op	    1056 B/op	       9 allocs/op
func Decode(s string, n int) string {
	result := make([]byte, len(s))

	for i, pos := range getKey(n, len(s)) {
		result[pos] = s[i]
	}

	return string(result)
}

func getKey(n int, l int) []int {
	result := make([]int, 0, l)
	step := n - 1

	for i := 0; i < n; i++ {
		for j, d := 0, 1; j < l; j, d = j+step, -d {
			if i == 0 { // top
				if d > 0 {
					result = append(result, j)
				}
				continue
			}

			if i == n-1 { // bottom
				if d < 0 {
					result = append(result, j)
				}
				continue
			}

			if d > 0 && j+i < l {
				result = append(result, j+i)
			}

			if d < 0 && step-i+j < l {
				result = append(result, step-i+j)
			}
		}
	}

	return result
}

func getKeySlow(n int, l int) []int {
	mx := make([][]int, n)
	for i := 0; i < n; i++ {
		mx[i] = make([]int, 0)
	}

	d := 1
	for i, j := 0, 0; i < l; i++ {
		mx[j] = append(mx[j], i)
		j += d
		if j == n-1 || j == 0 {
			d = -d
		}
	}

	result := make([]int, 0, l)
	for i := 0; i < n; i++ {
		result = append(result, mx[i]...)
	}

	return result
}
