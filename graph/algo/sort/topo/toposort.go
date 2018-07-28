package topo

import "github.com/go-li/libgeneration/graph"
import heap "github.com/go-li/libgeneration/heap/tree"

func cmp_int(a, b *int) int {
	return *a - *b
}

type intRoot struct {
	full_heaps_tree *intNode
	working_heap    []int
	length          int
}

type intNode struct {
	order_link [2]*intNode
	full_heap  []int
}

func Sort(api graph.Graph, graph_data *) (out []int) {

	var incoming []int

	var visited_nodes = 0

	var root intRoot

	graph.ForeachEdge(api, graph_data, func(i int, j int, d float64) {
//			print("from ")
//			print(i)
//			print(" to ")
//			print(j)
//			println("")

			for j >= len(incoming) {
				incoming = append(incoming, 0)
			}
			for i >= len(incoming) {
				incoming = append(incoming, 0)
			}

			incoming[j]++


	})

//	println("--")

	for i := range incoming {
		if incoming[i] == 0 {
			var sli = [1]int{i}

//			print(" adding ")
//			println(i)
			heap.Insert(&root, sli[0:], cmp_int)
		}
	}
//	println("--")
	for {

		var ptr = heap.Pop(&root, cmp_int)

		if ptr == nil {
			return out
		}

//		print("popped ")

		var u = *ptr

//		println(u)

		out = append(out, u)


		graph.ForeachEdgeFrom(u, api, graph_data, func(_ int, j int, d float64) {

//			print("from ")
//			print(u)
//			print(" to ")
//			print(j)
//			println("")

			incoming[j]--

			if incoming[j] == 0 {

				var sli = [1]int{j}

//				print(" adding ")
//				println(j)

				heap.Insert(&root, sli[0:], cmp_int)
			}


		})


		visited_nodes++

	}

	if visited_nodes != len(incoming) {
		return nil
	}

	return out
}
