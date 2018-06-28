package clusterer

import (
	"math"
	"sync"
)

/*
Glossary

N - power of set of elements to clustering
Each element represents as it's index 0...N-1

DistanceMatrix[i][j] - oriented distance between i-th and j-th elements
centroids[i] - index of element used as centroid for i-th cluster
clustered[i] - index of current centroid for i-th element
clusters[i][j] - index of j-th element in i-th cluster
*/

// GrabElementsToClusters returns list of indexes. Clustered[i] is index of element used as centroid for cluster containing i-th element
func MakeClusters(dm DistanceMatrix, centroids []int, count int) Clusters {
	clusters := NewClusters(count)
	//var wg sync.WaitGroup
	l := len(dm)
	for i := 0; i < l; i++ { // for each point
		//wg.Add(1)
		//go func(i int) {
		//defer wg.Done()
		min := math.MaxFloat64
		row := dm[i]
		var cn int // index of nearest centroid for i-th point
		// find nearest centroid
		for j, c := range centroids { // c is index of element marked as centroid
			if row[c] < min { //distance between point and centroid
				cn, min = j, row[c]
			}
		}
		clusters[cn].Add(i) // mark i-th point as clustered to cluster using cn-th element as centroid
		// In different words, cn-th element is centroid for cluster containing i-th point
		//}(i)
	}
	//wg.Wait()
	return clusters
}

// Recalculate centroids returns next Centroids
// Next centroid for cluster is an element in cluster minimising sum to all other elements in this cluster
func RecalculateCentroids(dm DistanceMatrix, clusters Clusters, count int, findCentroid FindCentroidFunc) Centroids {
	newCentroids := NewCentroids(count)
	var wg sync.WaitGroup
	for i, cluster := range clusters {
		wg.Add(1)
		go func (i int, cluster Cluster) {
			defer wg.Done()
			newCentroids[i] = findCentroid(dm, cluster)
		}(i, cluster)
	}
	wg.Wait()
	return newCentroids
}

func ChangesCount(c1, c2 Centroids, count int) int {
	changes := 0
	for i := 0; i < count; i++ {
		if c1[i] != c2[i] {
			changes++
		}
	}
	return changes
}

// ClusterizeKMeans
func ClusterizeKMeans(dm DistanceMatrix, count int, initCentroids InitCentroidsFunc, findCentroid FindCentroidFunc) Clusters {
	centroids := initCentroids(dm, count) // indexes of centroid points
	for {
		clusters := MakeClusters(dm, centroids, count)
		newCentroids := RecalculateCentroids(dm, clusters, count, findCentroid)
		if changed := ChangesCount(centroids, newCentroids, count); changed == 0 {
			return clusters
		} else {
			centroids = newCentroids
		}
	}
}
