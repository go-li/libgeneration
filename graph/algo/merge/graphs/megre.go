package graphs

import "github.com/go-li/libgeneration/graph"

func Merge(api graph.Graph, graphs_datas []*) (merged_graph_data *) {

	if len(graphs_datas) == 0 {
		return nil
	}

	merged_graph_data = api.New(graphs_datas[0])


	for g := range graphs_datas {


		graph.ForeachEdge(api, graphs_datas[g], func(i int, j int, d float64) {


			api.Set(i, j, d, merged_graph_data)

		})
	}

	api.Modify(merged_graph_data)

	return merged_graph_data
}
