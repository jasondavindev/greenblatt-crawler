package greenblatt

import (
	"sort"

	"github.com/jasondavindev/statusinvest-crawler/requester_types"
	"github.com/oleiade/reflections"
)

// SortAscBy Ordem crescente
func SortAscBy(d requester_types.StatusInvestResponse, name string) requester_types.StatusInvestResponse {
	sort.SliceStable(d, func(i, j int) bool {
		fpl := GetOrZero(reflections.GetField(&d[i], name))
		spl := GetOrZero(reflections.GetField(&d[j], name))

		return fpl < spl
	})

	return d
}

// SortDescBy Ordem decrescente
func SortDescBy(d requester_types.StatusInvestResponse, name string) requester_types.StatusInvestResponse {
	sort.SliceStable(d, func(i, j int) bool {
		fpl := GetOrZero(reflections.GetField(&d[i], name))
		spl := GetOrZero(reflections.GetField(&d[j], name))

		return fpl > spl
	})

	return d
}

// FindCompanyByName Busca empresa por nome
func FindCompanyByName(d requester_types.StatusInvestResponse, name string) int {
	for i := range d {
		if d[i].Ticker == name {
			return i
		}
	}

	return -1
}

// SortCompanies Ordena as companhias pelos ranks
func SortCompanies(d requester_types.StatusInvestResponse) map[string]RankItem {
	m := make(map[string]RankItem)

	SetRankNameBySlice(&m, SortAscBy(d, "PL"), "PlPosition")
	SetRankNameBySlice(&m, SortDescBy(d, "Roic"), "RoicPosition")
	SetRankNameBySlice(&m, SortDescBy(d, "Dividendo"), "DyPosition")
	CalculateFinalRankField(&m)
	return m
}

// GetSortedByFinalPosition Ordena companhias pelo rank final
func GetSortedByFinalPosition(m map[string]RankItem, d requester_types.StatusInvestResponse) requester_types.StatusInvestResponse {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for k, v := range m {
		ss = append(ss, kv{Key: k, Value: v.FinalPosition})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	s := requester_types.StatusInvestResponse{}

	for _, kv := range ss {
		idx := FindCompanyByName(d, kv.Key)
		c := d[idx]
		s = append(s, c)
	}

	return s
}
