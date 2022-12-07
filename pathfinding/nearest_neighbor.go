package pathfinding

import (
	"math/rand"

	"github.com/Murchik/collector_routes/models"
)

// Удалить из массива s первый элемент с значением v
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

// Входные данные - список банкоматов, матрица расстояний, начальный банкомат
// Выходные данные - маршрут
func Pathfinding(ATMS []models.ATM, matr [][]float64, start models.ATM) []int {

	var unvisited_nodes []int  // Непосещённые вершины (храним Id банкоматов)
	var visited_nodes []int    // Посещенные вершины (храним Id банкоматов)
	var current_node int       // Текущий банкомат
	var time float64           // Текущая продолжительность рабочего дня
	var max_time float64 = 175 // Максимальная продолжительность рабочего дня

	// Отмечаем все банкоматы как непосещенные
	for i := 0; i < len(ATMS); i++ {
		unvisited_nodes = append(unvisited_nodes, ATMS[i].Id)
	}

	visited_nodes = append(visited_nodes, start.Id)              // Отмечаем стартовый посещенным
	unvisited_nodes = delete_by_value(unvisited_nodes, start.Id) // Удаляем банкомат из непосещенных
	current_node = visited_nodes[0]                              // Id последнего посещенного банкомата

	for time < max_time {
		// Если укладываемся в максимальную продолжительность рабочего дня, то увеличиваем текущую
		if len(unvisited_nodes) == 0 {
			break
		}

		nearest_distance := matr[current_node][unvisited_nodes[0]] // расстояние до ближайшего банкомата
		nearest_i := unvisited_nodes[0]                            // Ближайший банкомат к последнему посещенному (Id банкомата)

		// Если осталась непосещенной один банкомат
		if len(unvisited_nodes) == 1 {
			temp := matr[current_node][unvisited_nodes[0]]
			// Если укладываемся в макс. время, то увеличиваем текущую
			if time+temp+matr[unvisited_nodes[0]][start.Id] < max_time {
				time += temp
				visited_nodes = append(visited_nodes, nearest_i)
				delete_by_value(unvisited_nodes, nearest_i)
			}
			break
		}

		// Ищем ближайший банкомат
		for i := 1; i < len(unvisited_nodes); i++ {
			temp := matr[current_node][unvisited_nodes[i]]
			if temp < nearest_distance {
				if time+temp+matr[i][start.Id] < max_time {
					nearest_distance = temp
					nearest_i = unvisited_nodes[i]
				}
			}
		}

		// Если укладываемся в максимальную продолжительность рабочего дня, то увеличиваем текущую
		if time+matr[current_node][nearest_i]+matr[nearest_i][start.Id] <= max_time {
			time += matr[current_node][nearest_i]
			current_node = nearest_i
			visited_nodes = append(visited_nodes, current_node)              // Отмечаем банкомат посещенным
			unvisited_nodes = delete_by_value(unvisited_nodes, current_node) // Удаляем банкомат из непосещенных
		} else {
			break
		}

	}

	// Возвращаемся в начальный банкомат
	visited_nodes = append(visited_nodes, start.Id)
	time += matr[current_node][start.Id]

	return visited_nodes
}

// Заполняем матрицу расстояний случайными данными
func CreateDistanceMatrix(qnt int) [][]float64 {

	arr := make([][]float64, qnt)

	for i := range arr {
		arr[i] = make([]float64, qnt)
	}

	for i := 0; i < qnt; i++ {
		for j := 0; j < qnt; j++ {
			if i == j {
				arr[i][j] = 0
			} else {
				if j < i {
					arr[i][j] = arr[j][i]
				} else {
					arr[i][j] = rand.Float64()*30 + 30
				}
			}
		}
	}

	return arr
}

// Удаляем из массива банкоматов arr1 банкоматы с индексами, содержащимися в массиве arr2
func DeleteAtmsFromArray(arr1 []models.ATM, arr2 []int) []models.ATM {
	for i := 0; i < len(arr1); i++ {
		for j := 1; j < len(arr2)-1; j++ {
			if arr1[i].Id == arr2[j] {
				arr1 = append(arr1[:i], arr1[i+1:]...)
				i -= 1
				break
			}
		}
	}

	return arr1
}
