package clusterer

import (
	"math/rand"
	"time"
	"math"
	"sync"
)

type InitCentroidsFunc func(DistanceMatrix, int) Centroids

func kmppFirstCentroidRandom(dm DistanceMatrix) int {
	l := len(dm)
	return rand.Intn(l)
}

func kmppFirstCentroidMinDist(dm DistanceMatrix) int {
	l := len(dm)
	dist := make([]float64, l)
	var wg sync.WaitGroup
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < l; j++ {
				dist[i] += dm[i][j]
			}
		}(i)
	}
	wg.Wait()
	min := math.MaxFloat64
	var index int
	for i := 0; i < l; i++ {
		if dist[i] < min {
			min = dist[i]
			index = i
		}
	}
	return index
}

func InitCentroidsKMPPRand(dm DistanceMatrix, count int) Centroids {
	return initCentroidsKMPP(dm, count, kmppFirstCentroidRandom)
}

func InitCentroidsKMPPDet(dm DistanceMatrix, count int) Centroids {
	return initCentroidsKMPP(dm, count, kmppFirstCentroidMinDist)
}

func initCentroidsKMPP(dm DistanceMatrix, count int, first func(DistanceMatrix) int) Centroids {
	rand.Seed(time.Now().Unix())
	l := len(dm)
	centroids := NewCentroids(count)
	centroids[0] = first(dm)
	//
	current := 1
	for current < count{
		distances := make([]float64, l)
		sdx := 0.0
		var wg sync.WaitGroup
		for i := 0; i < l; i++ { // for each point
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				min := math.MaxFloat64
				for _, c := range centroids { // find distance to nearest centroid
					min = math.Min(min, dm[i][c])
				}
				distances[i] = min

			}(i)
		}
		wg.Wait()
		for i := 0; i < l; i++ {
			sdx += distances[i]
		}

		rnd := rand.Float64() * sdx
		sdx2 := 0.0
		for i, dist := range distances {
			sdx2 += dist
			if sdx2 > rnd {
				centroids[current] = i
				current++
				break
			}
		}
	}
	//
	return centroids
}

func InitCentroidsFirstN(dm DistanceMatrix, count int) Centroids {
	centroids := NewCentroids(count)
	for i := 0; i < count; i++ {
		centroids[i] = i
	}
	return centroids
}

func InitCentroidsFibN(dm DistanceMatrix, count int) Centroids {
	centroids := NewCentroids(count)
	a, b := 0, 1
	for i := 0; i < count; i++ {
		centroids[i] = a % count
		a, b = b, (a + b) % count
	}
	return centroids
}

func InitCentroidsRandom(dm DistanceMatrix, count int) Centroids {
	rand.Seed(time.Now().Unix())
	l := len(dm)
	centroids := NewCentroids(count)
	current := 0
	for current < count {
		for {
			next := rand.Intn(l)
			if indexOf(next, centroids) == -1 {
				centroids[current] = next
				current++
				break
			}
		}
	}
	return centroids
}
