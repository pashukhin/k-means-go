package clusterer

type Points []Point

func NewPoints(count int) Points {
	return make(Points, count)
}

func (p *Points) Add(point Point) { // use pointer receiver is very important because this method changes receiver like slice = append(slice, elem)
	*p = append(*p, point)
}

func (p Points) Get(i int) Point {
	return p[i]
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Points) Less(i, j int) bool {
	d := min_int(p[i].GetDimension(), p[j].GetDimension())
	for k := 0; k < d; k++ {
		if p[i].GetCoord(k) != p[j].GetCoord(k) {
			return p[i].GetCoord(k) < p[j].GetCoord(k)
		}
	}
	return false
}

func (p Points) Len() int {
	return len(p)
}

