package accumulate

func Accumulate(a []string, f func(string) string) []string {
	result := make([]string, len(a))
	for i, item := range a {
		result[i] = f(item)
	}
	return result
}
