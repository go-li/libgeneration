package graph

import "math"
import "github.com/go-li/libgeneration/sort/quick"

type SparseGraph struct{}

type DataSparseGraph = [3]map[int][]int

func (sg SparseGraph) New(graph *DataSparseGraph) *DataSparseGraph {
	if graph == nil {
		graph = &DataSparseGraph{make(map[int][]int),
			make(map[int][]int), make(map[int][]int)}
	}
	return graph
}

func (sg SparseGraph) Get(from, to int, graph *DataSparseGraph) (float64, int) {
	const intmin = -int(^uint(0)>>1) - 1

	edges := graph[0][from]
	for i := range edges {
		if edges[i] == to {
			if i+1 >= len(edges) {
				return 1.0, intmin
			}
			return 1.0, edges[i+1]
		}
		if edges[i] > to {
			return math.Inf(1), edges[i]
		}
	}
	return math.Inf(1), intmin
}

func (sg SparseGraph) Set(from, to int, edge float64, graph *DataSparseGraph) {
	if math.IsNaN(edge) || math.IsInf(edge, 0) {
		graph[2][from] = append(graph[2][from], to)
	} else {
		graph[1][from] = append(graph[1][from], to)
	}
}

func (sg SparseGraph) Modify(graph *DataSparseGraph) {

	for n, s := range graph[0] {

		deletion := graph[1][n]

		quick.Sort(deletion, cmp_ints)

		graph[1][n] = deletion

		s = sorted_sorted_delete(s, deletion, cmp_ints)

		var temp int
		adaptive_sort(s, &temp, cmp_ints)

		deletion = graph[2][n]

		quick.Sort(deletion, cmp_ints)

		s = sorted_sorted_delete(s, deletion, cmp_ints)

		graph[0][n] = s
	}

	for n, s := range graph[1] {

		var lg = len(graph[0][n])

		graph[0][n] = append(graph[0][n], s...)

		if len(s) != 0 && lg != 0 {
			quick.Sort(graph[0][n], cmp_ints)
		}

	}
	graph[1] = make(map[int][]int)
	graph[2] = make(map[int][]int)

}
