package clusterer

import (
	"math"
)

// EuclideanMetric ...
type EuclideanMetric struct{}

// GetDistance - return euclidean distance between two points
func (e *EuclideanMetric) GetDistance(p1 IPoint, p2 IPoint) float64 {
	r := 0.0
	l := min_int(p1.GetDimension(), p2.GetDimension())
	for i := 0; i < l; i++ {
		r += sqr(p1.GetCoord(i) - p2.GetCoord(i))
	}
	return math.Sqrt(r)
}
