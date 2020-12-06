package main

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/csv"
	"github.com/jasondavindev/statusinvest-crawler/greenblatt"
	"github.com/jasondavindev/statusinvest-crawler/requester"
)

func main() {
	filters := config.CfgFactory("filtros.yml")
	ps := requester.CreateRequestParams(filters)

	r := requester.Find(ps)
	m := greenblatt.SortCompanies(r)
	fp := greenblatt.GetSortedByFinalPosition(m, r)
	csv.Write(fp)
}
