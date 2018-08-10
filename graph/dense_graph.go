package graph

import "math"

type DenseGraph struct{}

type ApiDenseGraph interface {
	New(*DataDenseGraph) *DataDenseGraph
	Get(int, int, *DataDenseGraph) (float64, int)
	Set(int, int, float64, *DataDenseGraph)
	Modify(*DataDenseGraph)
	Each(*DataDenseGraph, func(int))
	Transpose(*DataDenseGraph)
}

type DataDenseGraph = struct {
	upper    []byte
	diagonal []byte
	lower    []byte
	adds     map[[2]uint32]struct{}
	dels     map[[2]uint32]struct{}
}

func (sg DenseGraph) Transpose(m *DataDenseGraph) {
	m.lower, m.upper = m.upper, m.lower
}

func DenseGraphUndirectify1(m *DataDenseGraph) *DataDenseGraph {
	m.lower = m.upper
	return m
}

func DenseGraphUndirectify2(m *DataDenseGraph) *DataDenseGraph {
	m.upper = m.lower
	return m
}

func DenseGraphDirectify1(m *DataDenseGraph) *DataDenseGraph {
	copy(m.lower, m.upper)
	return m
}

func DenseGraphDirectify2(m *DataDenseGraph) *DataDenseGraph {
	copy(m.upper, m.lower)
	return m
}

func DenseGraphUnloopify(m *DataDenseGraph) *DataDenseGraph {
	m.diagonal = nil
	return m
}

func DenseGraphDirected(m *DataDenseGraph) bool {
	if m == nil {
		return true
	}
	if m.upper == nil {
		return true
	}
	if m.lower == nil {
		return true
	}
	if len(m.upper) == 0 {
		return true
	}
	if len(m.lower) == 0 {
		return true
	}
	return &m.upper[0] != &m.lower[0]
}

func DenseGraphUndirected(m *DataDenseGraph) bool {
	if m == nil {
		return true
	}
	if m.upper == nil {
		return true
	}
	if m.lower == nil {
		return true
	}
	if len(m.upper) == 0 {
		return true
	}
	if len(m.lower) == 0 {
		return true
	}
	return &m.upper[0] == &m.lower[0]
}

func DenseGraphLen(m *DataDenseGraph) (l int) {

	mat := len(m.lower)

	if len(m.upper) > len(m.lower) {
		mat = len(m.upper)
	}

	l = int(0.5*math.Sqrt(1+8*float64(mat))) + 1

	if len(m.diagonal) > l {
		return len(m.diagonal)
	}

	return l
}

func (sg DenseGraph) New(graph *DataDenseGraph) *DataDenseGraph {
	if graph == nil {
		graph = &DataDenseGraph{nil, nil, nil,
			make(map[[2]uint32]struct{}),
			make(map[[2]uint32]struct{})}
	}
	return graph
}

func (sg DenseGraph) Set(from, to int, edge float64, graph *DataDenseGraph) {

	if math.IsNaN(edge) || math.IsInf(edge, 0) {
		graph.dels[[2]uint32{uint32(from), uint32(to)}] = struct{}{}
	} else {
		graph.adds[[2]uint32{uint32(from), uint32(to)}] = struct{}{}
	}
}

func set(small, large int, data []byte, item byte) []byte {

	index := small + (large*(large-1))/2

	if item > 0 {

		for index >= len(data) {
			data = append(data, 0)
		}

		data[index] = item

	} else {

		if index < len(data) {
			data[index] = 0
		}

		if index+1 == len(data) {
			for (len(data) > 0) && (data[len(data)-1] == 0) {
				l := len(data) - 1
				data = data[0:l]
			}

		}

	}

	return data
}

func fetch(m *DataDenseGraph, lower []byte, upper []byte, y int, to int) (f float64, i int) {
	const intmin = -int(^uint(0)>>1) - 1

	var n = y - 1
	n = (n*n + n) / 2

	if y > to {
		if (n+to < len(lower)) && (lower[n+to] != 0) {
			f = 1.0
		} else {
			f = math.Inf(1)
		}
	} else if y == to {
		if (len(m.diagonal) > y) && (m.diagonal[y] != 0) {
			f = 1.0
		} else {
			f = math.Inf(1)
		}
	} else {
		var t = to - 1
		t = (t*t + t) / 2
		if (t+y < len(upper)) && (upper[t+y] != 0) {
			f = 1.0
		} else {
			f = math.Inf(1)
		}
	}

	for j := n + to + 1; j < (n+y) && j < len(lower); j++ {
		if lower[j] != 0 {

			return f, j - n

		}
	}
	if (y > to) && (len(m.diagonal) > 0) {
		if m.diagonal[y] != 0 {

			return f, y

		}
	}

	var x = y
	var j = 0
	if y < to {

		j = (to*to + to) / 2
		x += to

	} else {

		j = (y*y + y) / 2

	}

	for ; j+y < len(upper); j += x {

		if upper[j+y] != 0 {
			return f, x + 1
		}
		x++
	}

	return f, intmin
}

func (sg DenseGraph) Get(from int, to int, graph *DataDenseGraph) (float64, int) {
	return fetch(graph, graph.lower, graph.upper, from, to)
}

func (sg DenseGraph) Modify(graph *DataDenseGraph) {
	for v := range graph.dels {

		if v[0] == v[1] {
			if int(v[0]) >= len(graph.diagonal) {
				continue
			}
			graph.diagonal[v[0]] = 0

			if int(v[0])+1 == len(graph.diagonal) {
				for graph.diagonal[len(graph.diagonal)-1] == 0 {
					l := len(graph.diagonal) - 1
					graph.diagonal = graph.diagonal[0:l]
				}
			}
			continue
		}

		if v[0] < v[1] {
			graph.upper = set(int(v[0]), int(v[1]), graph.upper, 0)
		}
		if v[0] > v[1] {
			graph.lower = set(int(v[1]), int(v[0]), graph.lower, 0)
		}

	}

	for v := range graph.adds {

		if v[0] == v[1] {
			for int(v[0]) >= len(graph.diagonal) {
				graph.diagonal = append(graph.diagonal, 0)
			}
			graph.diagonal[v[0]] = 0xff
			continue
		}

		if v[0] < v[1] {
			graph.upper = set(int(v[0]), int(v[1]), graph.upper, 0xff)
		}
		if v[0] > v[1] {
			graph.lower = set(int(v[1]), int(v[0]), graph.lower, 0xff)
		}

	}

	graph.dels = make(map[[2]uint32]struct{})
	graph.adds = make(map[[2]uint32]struct{})
}

func (sg DenseGraph) Each(graph *DataDenseGraph, callback func(int)) {
	for i := 0; i < DenseGraphLen(graph); i++ {
		callback(i)
	}
}
