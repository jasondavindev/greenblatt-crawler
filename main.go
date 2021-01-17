package main

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/csv"
	"github.com/jasondavindev/statusinvest-crawler/filter_classes"
	"github.com/jasondavindev/statusinvest-crawler/greenblatt"
	"github.com/jasondavindev/statusinvest-crawler/requester"
)

func main() {
	filters := config.CfgFactory("filtros.yml")
	addFilters := filter_classes.GetSectorsAddFilterable()
	r := requester.AsyncFind(&filters, addFilters)
	m := greenblatt.SortCompanies(r)
	fp := greenblatt.GetSortedByFinalPosition(m, r)
	csv.Write(fp)
}
