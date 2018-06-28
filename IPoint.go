package clusterer

type IPoint interface {
	GetCoord(index int) float64
	GetDimension() int
}
