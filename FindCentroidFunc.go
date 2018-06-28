package clusterer

import "math"

type FindCentroidFunc func (DistanceMatrix, Cluster) int

func FindCentroidMin (dm DistanceMatrix, cluster Cluster) int {
	var r int
	min := math.MaxFloat64
	for _, j := range cluster {
		sum := 0.0
		for _, k := range cluster {
			sum += dm[j][k] // distance between j-th and k-th elements
		}
		if sum < min {
			min = sum
			r = j
		}
	}
	return r
}

func FindCentroidCenter (dm DistanceMatrix, cluster Cluster) int {
	var r int
	min := math.MaxFloat64
	for _, j := range cluster {
		e := 0.0
		for _, k := range cluster {
			e = math.Max(e, dm[j][k])
		}
		if e <= min {
			min = e
			r = j
		}
	}
	return r
}
