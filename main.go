package main

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/csv"
	"github.com/jasondavindev/statusinvest-crawler/greenblatt"
	"github.com/jasondavindev/statusinvest-crawler/search"
)

func main() {
	configSetup := config.CfgFactory("filtros.yml")
	results := search.GetSearchMode(configSetup)
	sortedCompanies := greenblatt.SortCompanies(results)
	finalPosition := greenblatt.GetSortedByFinalPosition(sortedCompanies, results)
	csv.Write(finalPosition)
}
