package graph

type Graph interface {
	New(*) *
	Get(int, int, *) (float64, int)
	Set(int, int, float64, *)
	Modify(*)
}

func Dist(d float64, n int) float64 {
	return d
}

func Next(d float64, n int) int {
	return n
}
