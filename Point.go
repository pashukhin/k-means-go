package clusterer

// Point represents n-coordinate point
type Point []float64

// NewPoint makes a new point}
func NewPoint(coords ...float64) Point {
	p := Point{}
	for _, coord := range coords {
		p.add(coord)
	}
	return p
}

func (p *Point) add(coord float64) { // use pointer receiver is very important because this method changes receiver like slice = append(slice, elem)
	*p = append(*p, coord)
}

func (p Point) GetCoord(index int) float64 {
	if index < p.GetDimension() {
		return p[index]
	}
	return 0.0
}

func (p Point) GetDimension() int {
	return len(p)
}
