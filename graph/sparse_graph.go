package graph

import "github.com/go-li/libgeneration/dict/avl"
import "math"

type SparseGraph struct{}

type ApiSparseGraph interface {
	New(*DataSparseGraph) *DataSparseGraph
	Get(int, int, *DataSparseGraph) (float64, int)
	Set(int, int, float64, *DataSparseGraph)
	Modify(*DataSparseGraph)
	Each(*DataSparseGraph, func(int))
	Transpose(*DataSparseGraph)
}

func cmp0a(a, b *edge_node) int {
	return a.fromto[0] - b.fromto[1]
}
func cmp1a(a, b *edge_node) int {
	return a.fromto[1] - b.fromto[0]
}


func cmp0(a, b *edge_node) (o int) {
	o = a.fromto[0] - b.fromto[0]
	if o == 0 {
		o = a.fromto[1] - b.fromto[1]
	}
	return o
}

func cmp1(a, b *edge_node) (o int) {
	o = a.fromto[1] - b.fromto[1]
	if o == 0 {
		o = a.fromto[0] - b.fromto[0]
	}
	return o
}

func tree0(n *edge_node) (*[3]*edge_node, *byte) {
	return &n.link[0], &n.bal[0]
}

func tree1(n *edge_node) (*[3]*edge_node, *byte) {
	return &n.link[1], &n.bal[1]
}

type edge_node struct {
	link [2][3]*edge_node
	bal [2]byte
	fromto [2]int
}

type DataSparseGraph = struct {
	roots [2]*edge_node
	xor byte
	adds map[[2]int]struct{}
	dels map[[2]int]struct{}
}

func (sg SparseGraph) Transpose(d *DataSparseGraph) {
	d.xor ^= 1
}

func (sg SparseGraph) New(graph *DataSparseGraph) *DataSparseGraph {
	if graph == nil {
		graph = &DataSparseGraph{adds : make(map[[2]int]struct{}),
					dels : make(map[[2]int]struct{})}
	}
	return graph
}

func (sg SparseGraph) Get(from, to int, graph *DataSparseGraph) (float64, int) {
	const intmin = -int(^uint(0)>>1) - 1

	var node edge_node
	node.fromto[graph.xor] = from
	node.fromto[1^graph.xor] = to

	var cmps = [2]func(a, b *edge_node) (int) {cmp0, cmp1}
	var linkers = [2]func (n *edge_node) (*[3]*edge_node, *byte) {tree0, tree1}

	var found *edge_node = avl.Find(linkers[graph.xor], graph.roots[graph.xor], func(e *edge_node) byte {

		var delta = cmps[graph.xor](e, &node)
		if (delta == 0) {
			return 6
		}

		var r *edge_node = e.link[graph.xor][0]
		if r != nil {
			for r.link[graph.xor][1] != nil {
				r = r.link[graph.xor][1]
			}
		}

		if (delta >= 0 && r == nil) || (delta >= 0 && cmps[graph.xor](r, &node) < 0) {
			return 6
		}



		if (delta >= 0){
			return 0
		}

		return 1
	})
	if found == nil {
		return math.Inf(1), intmin
	}


	if (found.fromto[0] == node.fromto[0] && found.fromto[1] == node.fromto[1]) {
		var next *edge_node = avl.InOrderSuccessor(linkers[graph.xor], found)

		if next == nil || next.fromto[graph.xor] != node.fromto[graph.xor] {
			return 1.0, intmin
		} else {
			return 1.0, next.fromto[1^graph.xor]
		}
	}

	if node.fromto[graph.xor] == found.fromto[graph.xor] {
		return math.Inf(1), found.fromto[1^graph.xor]
	}

	return math.Inf(1), intmin
}

func (sg SparseGraph) Set(from, to int, edge float64, graph *DataSparseGraph) {
	if math.IsNaN(edge) || math.IsInf(edge, 0) {
		graph.dels[[2]int{from, to}] = struct{}{}
	} else {
		graph.adds[[2]int{from, to}] = struct{}{}
	}
}


func (sg SparseGraph) Modify(graph *DataSparseGraph) {

	var cmps = [2]func(a, b *edge_node) (int) {cmp0, cmp1}
	var linkers = [2]func (n *edge_node) (*[3]*edge_node, *byte) {tree0, tree1}

	for k := range graph.dels {
		var node edge_node
		node.fromto[graph.xor] = k[0]
		node.fromto[1^graph.xor] = k[1]

		var del1 *edge_node = avl.Delete(linkers[graph.xor], &(graph.roots[graph.xor]), &node, cmps[graph.xor])
		avl.Delete(linkers[1^graph.xor], &(graph.roots[1^graph.xor]), &node, cmps[1^graph.xor])

		if del1 == nil {
			print(k[0])
			print(" ")
			println(k[1])
		} else {

			del1.fromto[0] = -1
			del1.fromto[1] = -1
		}
	}

	for k := range graph.adds {
		var node edge_node
		node.fromto[graph.xor] = k[0]
		node.fromto[1^graph.xor] = k[1]

		var appended *edge_node = avl.Append(linkers[graph.xor], &(graph.roots[graph.xor]), &node, cmps[graph.xor])

		if appended != &node {
			appended.fromto = node.fromto
		}

		appended = avl.Append(linkers[1^graph.xor], &(graph.roots[1^graph.xor]), appended, cmps[1^graph.xor])
	}

	graph.dels = make(map[[2]int]struct{})
	graph.adds = make(map[[2]int]struct{})
}

func (sg SparseGraph) Each(graph *DataSparseGraph, callback func(int)) {
	const intmin = -int(^uint(0)>>1) - 1

	var cmps = [2]func(a, b *edge_node) (int) {cmp0a, cmp1a}
	var linkers = [2]func (n *edge_node) (*[3]*edge_node, *byte) {tree0, tree1}

	var node = intmin

	if graph.roots[0] == nil || graph.roots[1] == nil {
		return
	}

	var alpha *edge_node = avl.Find(linkers[graph.xor], graph.roots[graph.xor], nil)
	var beta *edge_node = avl.Find(linkers[1^graph.xor], graph.roots[1^graph.xor], nil)

	var cmp int

	for alpha != nil && beta != nil {

		cmp = cmps[graph.xor](alpha, beta)

		if cmp < 0 {
			if node < alpha.fromto[graph.xor] {
				node = alpha.fromto[graph.xor]
				callback(node)
			}

			alpha = avl.InOrderSuccessor(linkers[graph.xor], alpha)
		} else {
			if node < beta.fromto[1^graph.xor] {
				node = beta.fromto[1^graph.xor]
				callback(node)
			}

			beta = avl.InOrderSuccessor(linkers[1^graph.xor], beta)
		}
	}

	for alpha != nil {
		if node < alpha.fromto[graph.xor] {
			node = alpha.fromto[graph.xor]
			callback(node)
		}

		alpha = avl.InOrderSuccessor(linkers[graph.xor], alpha)
	}
	for beta != nil {
		if node < beta.fromto[1^graph.xor] {
			node = beta.fromto[1^graph.xor]
			callback(node)
		}

		beta = avl.InOrderSuccessor(linkers[1^graph.xor], beta)
	}

}
