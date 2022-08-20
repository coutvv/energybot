package main

import (
	"fmt"

	"github.com/coutvv/energybot/internal/energy/db"
	mappath "github.com/coutvv/energybot/internal/energy/mappath"
)

// Just example to use graph searching path in the db
func main() {
	var repository = db.NewSqliteRepository()

	cities := repository.GetAllCities()
	cables := repository.GetAllCables()
	fmt.Println(cables)
	graph :=	mappath.CablesToGraph(cables)
	fmt.Println(graph)
	for _, city := range cities {
		result := mappath.FindPathPrices(graph, city.Id)
		fmt.Println(result)
	}

}