package main

import (
	"fmt"
	"math"
)

type Edge struct {
	Node   string
	Weight int
}

type Graph = map[string][]Edge

func mainy() {
	graph := make(Graph)

	// Declared Graph
	graph["book"] = []Edge{{"lp", 5}, {"poster", 0}}
	graph["poster"] = []Edge{{"guitar", 30}, {"drums", 35}}
	graph["lp"] = []Edge{{"guitar", 15}, {"drums", 20}}
	graph["guitar"] = []Edge{{"piano", 20}}
	graph["drums"] = []Edge{{"piano", 10}}
	graph["piano"] = []Edge{}

	dijkstraTable := dijkstra(graph, "book")
	prettyPrint(dijkstraTable)
}

func prettyPrint(table DijkstraResultTable) {
	for k, v := range table {
		fmt.Printf("%s: %+v\n", k, v)
	}
}

type DijkstraAux struct {
	MinValue int
	Parent   string
	Checked  bool
}

const INF = math.MaxInt

type DijkstraResultTable = map[string]DijkstraAux

func dijkstra(graph Graph, start string) DijkstraResultTable {
	// tabela . vertice, valor minimo, parent
	aux := make(DijkstraResultTable)

	// Initialize Aux Table, all nodes are infinetly distant
	for node := range graph {
		aux[node] = DijkstraAux{INF, "", false}
	}
	aux[start] = DijkstraAux{0, "", true}

	currentCheapestNode := start
	cheapestCost := 0

	for currentCheapestNode != "" {
		auxEntry := aux[currentCheapestNode]
		auxEntry.Checked = true
		aux[currentCheapestNode] = auxEntry

		for _, edge := range graph[currentCheapestNode] {
			auxEntry := aux[edge.Node]

			if edge.Weight+cheapestCost < auxEntry.MinValue {
				auxEntry.MinValue = edge.Weight + cheapestCost
				auxEntry.Parent = currentCheapestNode
				aux[edge.Node] = auxEntry
			}
		}

		cheapestCost = math.MaxInt
		currentCheapestNode = ""

		for node, auxStruct := range aux {
			if !auxStruct.Checked && auxStruct.MinValue < cheapestCost {
				cheapestCost = auxStruct.MinValue
				currentCheapestNode = node
			}
		}
	}

	return aux
}
