package search

import "github.com/go-li/libgeneration/graph"

func Depth(from []int, visited map[int]struct{}, api graph.Graph, graph_data *, vertex func(int), edge func(int, int, float64)) {

	if visited == nil {
		visited = make(map[int]struct{})
	}

	for len(from) > 0 {

		at := from[0]

//		print("at ")
//		println(at)

		if vertex != nil {
			vertex(at)
		}

		visited[at] = struct{}{}
		from = from[1:]

		var add []int



		graph.ForeachEdgeFrom(at, api, graph_data, func(i int, j int, d float64) {

			if _, ok := visited[j]; ok {
				return
			}


//			print("taking edge ")
//			print(i)
//			print(" to ")
//			println(j)

			if edge != nil {
				edge(i,j,d)
			}

			add = append(add, j)


		})

		from = append(add, from...)

	}
}
