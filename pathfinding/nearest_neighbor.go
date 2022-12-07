package pathfinding

import (
	"github.com/Murchik/collector_routes/models"
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
func Pathfinding(ATMS []models.ATM, matr [][]float64, start models.ATM) []int {

	var unvisited_nodes []int // Непосещённые вершины храним Id банкоматов
	var visited_nodes []int   // Посещенные вершины храним Id банкоматов
	var current_node int
	var time float64
	var max_time float64 = 175 // Продолжительность рабочего дня

	// Отмечаем все банкоматы как непосещенные
	for i := 0; i < len(ATMS); i++ {
		unvisited_nodes = append(unvisited_nodes, ATMS[i].Id)
	}

	// Отмечаем стартовый посещенным
	visited_nodes = append(visited_nodes, start.Id)
	unvisited_nodes = delete_by_value(unvisited_nodes, start.Id)
	current_node = visited_nodes[0] // Id последнего посещенного банкомата

	for time < max_time {
		if len(unvisited_nodes) == 0 {
			break
		}

		nearest_distance := matr[current_node][unvisited_nodes[0]]
		nearest_i := unvisited_nodes[0] // Храним Id

		if len(unvisited_nodes) == 1 {
			temp := matr[current_node][unvisited_nodes[0]]
			if time+temp+matr[unvisited_nodes[0]][start.Id] < max_time {
				time += temp
			}
			break
		}

		for i := 1; i < len(unvisited_nodes); i++ {
			temp := matr[current_node][unvisited_nodes[i]]
			if temp < nearest_distance {
				if time+temp+matr[i][start.Id] < max_time {
					nearest_distance = temp
					nearest_i = unvisited_nodes[i]
				}
			}
		}

		if time+matr[current_node][nearest_i]+matr[nearest_i][start.Id] <= max_time {
			time += matr[current_node][nearest_i]
			current_node = nearest_i
			visited_nodes = append(visited_nodes, current_node)
			unvisited_nodes = delete_by_value(unvisited_nodes, current_node)
		} else {
			break
		}

	}

	visited_nodes = append(visited_nodes, start.Id)
	time += matr[current_node][start.Id]

	return visited_nodes
}
