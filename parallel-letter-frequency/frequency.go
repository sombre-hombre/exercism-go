package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	if len(texts) == 0 {
		return FreqMap{}
	}
	if len(texts) == 1 {
		return Frequency(texts[0])
	}

	var results = make(chan FreqMap)
	var f = func(t []string) {
		results <- ConcurrentFrequency(t)
	}

	m := len(texts) / 2
	go f(texts[:m])
	go f(texts[m:])

	r := <-results
	for k, v := range <-results {
		r[k] += v
	}

	return r
}
