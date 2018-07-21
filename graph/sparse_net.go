package graph

import "math"
import "github.com/go-li/libgeneration/sort/quick"

type adjacency = struct {
	target   int
	distance float64
}

func cmp_adj(a, b *adjacency) int {
	return a.target - b.target
}

type SparseNet struct{}

type DataSparseNet = [3]map[int][]adjacency

func (sg SparseNet) New(graph *DataSparseNet) *DataSparseNet {
	if graph == nil {
		graph = &DataSparseNet{make(map[int][]adjacency),
			make(map[int][]adjacency), make(map[int][]adjacency)}
	}
	return graph
}

func (sg SparseNet) Get(from, to int, graph *DataSparseNet) (float64, int) {
	const intmin = -int(^uint(0)>>1) - 1

	edges := graph[0][from]
	for i := range edges {
		if edges[i].target == to {
			if i+1 >= len(edges) {
				return edges[i].distance, intmin
			}
			return edges[i].distance, edges[i+1].target
		}
		if edges[i].target > to {
			return math.Inf(1), edges[i].target
		}
	}
	return math.Inf(1), intmin
}

func (sg SparseNet) Set(from, to int, edge float64, graph *DataSparseNet) {
	if math.IsNaN(edge) || math.IsInf(edge, 0) {
		graph[2][from] = append(graph[2][from], adjacency{to, edge})
	} else {
		graph[1][from] = append(graph[1][from], adjacency{to, edge})
	}
}

func (sg SparseNet) Modify(graph *DataSparseNet) {

	for n, s := range graph[0] {

		deletion := graph[1][n]

		quick.Sort(deletion, cmp_adj)

		graph[1][n] = deletion

		s = sorted_sorted_delete(s, deletion, cmp_adj)

		var temp adjacency
		adaptive_sort(s, &temp, cmp_adj)

		deletion = graph[2][n]

		quick.Sort(deletion, cmp_adj)

		s = sorted_sorted_delete(s, deletion, cmp_adj)

		graph[0][n] = s
	}

	for n, s := range graph[1] {

		var lg = len(graph[0][n])

		graph[0][n] = append(graph[0][n], s...)

		if len(s) != 0 && lg != 0 {
			quick.Sort(graph[0][n], cmp_adj)
		}

	}
	graph[1] = make(map[int][]adjacency)
	graph[2] = make(map[int][]adjacency)

}
