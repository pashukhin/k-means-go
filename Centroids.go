package clusterer

// Centroids : Centroids[i] is index of element which used as centroid for i-th cluster
type Centroids []int

func NewCentroids(count int) Centroids {
	return make(Centroids, count)
}
