package pascal

func Triangle(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, i+1)
		for j := range res[i] {
			res[i][j] = 1
		}
	}

	for i := range res {
		if i > 1 {
			for j := range res[i] {
				if j > 0 && j < len(res[i])-1 {
					res[i][j] = res[i-1][j] + res[i-1][j-1]
				}
			}
		}
	}

	return res
}
