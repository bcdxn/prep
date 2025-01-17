package ds

func ParseMatrixFromEdges[T comparable](edges [][]T) map[T][]T {
	matrix := make(map[T][]T)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]

		if _, ok := matrix[from]; !ok {
			matrix[from] = make([]T, 0)
		}

		matrix[from] = append(matrix[from], to)
	}

	return matrix
}
