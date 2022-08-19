package main

import (
	"fmt"

	"github.com/coutvv/energybot/internal/energy/db"
	mappath "github.com/coutvv/energybot/internal/energy/mappath"
)

// Just example to use graph searching path in the db
func main() {
	var repository = db.NewSqliteRepository()

	cables := repository.GetAllCables()
	fmt.Println(cables)
	graph :=	mappath.CablesToGraph(cables)
	fmt.Println(graph)
	result := mappath.FindPathPrices(graph, "1A")
	fmt.Println(result)

}