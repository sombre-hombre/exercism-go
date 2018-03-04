package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	result := make([]string, len(s)-n+1)
	for i := 0; i+n <= len(s); i++ {
		result[i] = s[i : i+n]
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return
	}

	return UnsafeFirst(n, s), true
}
