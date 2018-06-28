package clusterer

// Clusters : Clusters[i][j] is index of j-th element of i-th cluster
type Clusters []Cluster

func NewClusters(count int) Clusters {
	r := make(Clusters, count)
	for i := 0; i < count; i++ {
		r[i] = NewCluster()
	}
	return r
}
