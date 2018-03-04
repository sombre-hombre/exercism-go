package pascal

// Triangle computes Pascal's triangle up to a given number of rows.
func Triangle(n int) [][]int {
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if i == 0 {
				triangle[i][j] = 1
				continue
			}
			triangle[i][j] = get(triangle, i-1, j-1) + get(triangle, i-1, j)
		}
	}

	return triangle
}

func get(t [][]int, i, j int) int {
	if i < 0 || len(t) <= i {
		return 0
	}
	if j < 0 || len(t[i]) <= j {
		return 0
	}

	return t[i][j]
}
