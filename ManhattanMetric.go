package clusterer

import (
	"math"
)

// ManhattanMetric ...
type ManhattanMetric struct{}

// GetDistance - return euclidean distance between two points
func (e *ManhattanMetric) GetDistance(p1 IPoint, p2 IPoint) float64 {
	r := 0.0
	l := min_int(p1.GetDimension(), p2.GetDimension())
	for i := 0; i < l; i++ {
		r += math.Abs(p1.GetCoord(i) - p2.GetCoord(i))
	}
	return r
}
