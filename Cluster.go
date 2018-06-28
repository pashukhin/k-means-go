package clusterer

// Cluster - group of points
type Cluster []int

// NewCluster - makes new Cluster
func NewCluster() Cluster {
	return Cluster{}
}

func (c *Cluster) Add(i int) { // use pointer receiver is very important because this method changes receiver like slice = append(slice, elem)
	*c = append(*c, i)
}
