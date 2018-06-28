package clusterer

type IPoints interface {
	Swap(i, j int)
	Less(i, j int) bool
	Len() int
	Add(p IPoint)
	Get(i int) IPoint
}
