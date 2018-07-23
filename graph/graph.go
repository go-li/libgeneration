package graph

import "math"

type Graph interface {
	New(*) *
	Get(int, int, *) (float64, int)
	Set(int, int, float64, *)
	Modify(*)
	Each(*, func(int))
}

func Dist(d float64, n int) float64 {
	return d
}

func Next(d float64, n int) int {
	return n
}

func IsEdge(dist float64) bool {
	return !math.IsNaN(dist) && !math.IsInf(dist, 0)
}

func HasEdge(d float64, _ int) bool {
	return IsEdge(d)
}
