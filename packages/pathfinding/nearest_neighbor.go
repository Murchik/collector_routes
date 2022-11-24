package pathfinding

import (
	"github.com/Murchik/collector_routes/packages/atm"
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
func Pathfinding(ATMS []atm.ATM, matr [][]float64, start atm.ATM) []int {

	var unvisited_nodes []int // Непосещённые вершины храним Id банкоматов
	var visited_nodes []int   // Посещенные вершины храним Id банкоматов
	var current_node int
	var time float64
	var max_time float64 = 100 // Продолжительность рабочего дня

	// Отмечаем все банкоматы как непосещенные
	for i := 0; i < len(ATMS); i++ {
		unvisited_nodes = append(unvisited_nodes, ATMS[i].Id)
	}

	// Отмечаем стартовый посещенным
	visited_nodes = append(visited_nodes, start.Id)
	unvisited_nodes = delete_by_value(unvisited_nodes, start.Id)

	for time < max_time {
		if len(unvisited_nodes) == 0 {
			break
		}
		current_node = visited_nodes[0] // Id последнего посещенного банкомата
		nearest_distance := matr[current_node][unvisited_nodes[0]]
		nearest_i := unvisited_nodes[0] // Храним Id
		for i := 1; i < len(unvisited_nodes); i++ {
			temp := matr[current_node][unvisited_nodes[i]] // возможно где-то путаница c Id и порядковым номером
			if temp < nearest_distance {
				if temp+matr[i][start.Id] < max_time {
					nearest_distance = temp
					nearest_i = unvisited_nodes[i]
				}
			}
		}
		if nearest_i == current_node {
			break
		} else {
			time += matr[current_node][nearest_i]
			current_node = nearest_i
			visited_nodes = append(visited_nodes, current_node)
			unvisited_nodes = delete_by_value(unvisited_nodes, current_node)
		}
	}

	visited_nodes = append(visited_nodes, start.Id)

	return visited_nodes
}
