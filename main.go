package main

import (
	"github.com/jasondavindev/statusinvest-crawler/csv"
	"github.com/jasondavindev/statusinvest-crawler/greenblatt"
	"github.com/jasondavindev/statusinvest-crawler/requester"
)

func main() {
	t := requester.InputItems{
		Dividendo:    []float32{1, 25},
		PL:           []float32{0.1, 20},
		DivLiqEbit:   []float32{0, 20},
		DivLiqPatLiq: []float32{0, 20},
		Roe:          []float32{0.1, 100},
		Roic:         []float32{0.1, 100},
		Roa:          []float32{0.1, 100},
		LiqMedDiaria: []float32{0, 10000000},
	}

	ps := requester.CreateRequestParams(t)

	r := requester.Find(ps)
	m := greenblatt.SortCompanies(r)
	fp := greenblatt.GetSortedByFinalPosition(m, r)
	csv.Write(fp)
}
