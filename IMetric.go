package clusterer

// IMetric - interface for metrics
type IMetric interface {
	GetDistance(p1 IPoint, p2 IPoint) float64
}
