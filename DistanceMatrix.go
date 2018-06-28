package clusterer

// DistanceMatrix - represents all distances between given Points
type DistanceMatrixRow []float64

func NewDistanceMatrixRow(count int) DistanceMatrixRow {
	return make(DistanceMatrixRow, count)
}

type DistanceMatrix []DistanceMatrixRow

func NewDistanceMatrix(count int) DistanceMatrix {
	r := make(DistanceMatrix, count)
	for i := 0; i < count; i++ {
		r[i] = NewDistanceMatrixRow(count)
	}
	return r
}
