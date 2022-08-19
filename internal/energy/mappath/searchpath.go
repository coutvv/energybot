package mappath

import "github.com/coutvv/energybot/internal/energy/db/entity"

func FindPathPrices(graph map[string][]Path, start string) map[string]Path{
	checked := make(map[string]bool)
	prices := make(map[string]Path)
	prices[start] = Path{name:start, weight:0}

	queue := []string{start}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		checked[item] = true
		itemPrice := prices[item]
		for _, kid := range graph[item] {
			if !checked[kid.name]  {
				queue = append(queue, kid.name)
			}
			val, ok := prices[kid.name]
			newWeight := kid.weight + itemPrice.weight
			if !ok {
				prices[kid.name] = Path{name: item, weight: newWeight}
			} else {
				if val.weight > newWeight {
					checked[kid.name] = false // check again
					queue = append(queue, kid.name)
					prices[kid.name] = Path{name: item, weight: newWeight}
				}
			}
		}
	}
	return prices
}

type Path struct {
	name string
	weight int
}

func NewPath(name string, weight int) Path {
	return Path{name: name, weight: weight}
}


func CablesToGraph(cables []entity.Cable) map[string][]Path {
	result := make(map[string][]Path)

	for _, cable := range cables {
		addCable(result, cable.Src, cable.Dest, cable.Price)
		addCable(result, cable.Dest, cable.Src, cable.Price)
	}

	return result
}

func addCable(graph map[string][]Path, src string, dest string, price int) {
		cityPaths, ok := graph[src]
		if !ok {
			graph[src] = []Path{{name: dest, weight: price}}
		} else {
			graph[src] = append(cityPaths, Path{name: dest, weight: price})
		}
}