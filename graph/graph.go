package graph

import "math"

type Graph interface {
	New(*) *
	Get(int, int, *) (float64, int)
	Set(int, int, float64, *)
	Modify(*)
	Each(*, func(int))
	Transpose(*)
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

func ForeachEdge(g Graph, graph_data *, callback func (int, int, float64)) {

	g.Each(graph_data, func(i int) {
		const intmin = -int(^uint(0)>>1) - 1
		var j = 0
		var j2 = 0
		var d float64

		for {
			d, j2 = g.Get(i, j, graph_data)
			if !IsEdge(d) {
				j = j2
				if j2 < 0 {
					break
				}
				continue
			}

			callback(i, j, d)

			if j2 < 0 {
				break
			}

			j = j2
		}
	})

}

func ForeachEdgeFrom(i int, g Graph, graph_data *, callback func (int, int, float64)) {
	const intmin = -int(^uint(0)>>1) - 1

	var j = 0
	var j2 = 0
	var d float64
	for {
		d, j2 = g.Get(i, j, graph_data)
		if !IsEdge(d) {
			j = j2
			if j2 < 0 {
				break
			}
			continue
		}

		callback(i, j, d)

		if j2 < 0 {
			break
		}

		j = j2
	}
}
