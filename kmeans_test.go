package clusterer

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestClusterizeKMeans(t *testing.T) {

	maxx, maxy := 100, 100
	start, points, l, elapsed, dm := InitTesting(maxx, maxy)

	count := 500
	start = time.Now()
	clusters := ClusterizeKMeans(dm, count, InitCentroidsKMPPDet, FindCentroidCenter)
	elapsed = time.Now().Sub(start)
	fmt.Println("clusterize", l, count, elapsed)

	start = time.Now()
	vis(clusters, nil, points, "KMeans")
	elapsed = time.Now().Sub(start)
	fmt.Println("vis", elapsed)
}

func InitTesting(maxx int, maxy int) (time.Time, []IPoint, int, time.Duration, DistanceMatrix) {

	var points []IPoint
	var l int
	TimeDecorator(func() []interface{}{
		points = initRegularPoints(maxx,maxy)
		l = len(points)
		return []interface{}{"init points", l}
	})
	metric := EuclideanMetric{}
	//metric := ManhattanMetric{}
	start := time.Now()
	dm := NewDistanceMatrix(l)
	elapsed := time.Now().Sub(start)
	fmt.Println("create distance matrix", l, elapsed)
	start = time.Now()
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < l; j++ {
				dm[i][j] = metric.GetDistance(points[i], points[j])
			}
		}(i)
	}
	wg.Wait()
	elapsed = time.Now().Sub(start)
	fmt.Println("fill distance matrix", l*l, elapsed)
	return start, points, l, elapsed, dm
}
