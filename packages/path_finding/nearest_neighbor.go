package pathfinding

import (
	atm "github.com/Murchik/collector_routes/packages/ATM"
)

func delete_by_value(s []int, v int) []int {
	var index int
	for i, value := range s {
		if value == v {
			index = i
			break
		}
	}
	s = append(s[:index], s[index+1:]...)
	return s
}

// Входные  параметры - список банкоматов, матрица смежности, начальный банкомат
func pathfinding(ATMS []atm.ATM, matr [][]float64, start atm.ATM) []int {

	var unvisited_nodes []int // Непосещённые вершины
	var visited_nodes []int   // Посещенные вершины
	var current_node int
	var time float64
	var max_time float64 = 480

	for i := 0; i < len(ATMS); i++ {
		unvisited_nodes = append(unvisited_nodes, ATMS[i].id)
	}

	visited_nodes = append(visited_nodes, start.id)
	unvisited_nodes = delete_by_value(unvisited_nodes, start.id)

	for time < max_time || len(unvisited_nodes) != 0 {

		current_node = visited_nodes[0]
		nearest_distance := matr[current_node][unvisited_nodes[1]]
		nearest_i := 1

		for i := 1; i < len(unvisited_nodes); i++ {

			temp := matr[current_node][unvisited_nodes[i]]

			if temp < nearest_distance {
				if temp+matr[i][start.id] < float64(max_time) {
					nearest_distance = temp
					nearest_i = i
				}
			}
		}
		if nearest_i == 1 {
			break
		} else {
			time += matr[current_node][unvisited_nodes[nearest_i]]
			current_node = unvisited_nodes[nearest_i]
			visited_nodes = append(visited_nodes, current_node)
			unvisited_nodes = delete_by_value(unvisited_nodes, current_node)
		}

	}

	return visited_nodes
}
