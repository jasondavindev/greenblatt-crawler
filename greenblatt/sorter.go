package greenblatt

import (
	"log"
	"sort"

	"github.com/jasondavindev/statusinvest-crawler/requester"
	"github.com/oleiade/reflections"
)

// GreenblattRank Estrutura de rankeamento
type GreenblattRank struct {
	PlPosition    int
	RoicPosition  int
	DyPosition    int
	FinalPosition int
}

// GetOrZero Retorna o valor ou zero
func GetOrZero(v interface{}, e interface{}) float64 {
	if e != nil {
		log.Fatal(e)
		return 0
	}

	if x, ok := v.(float64); ok {
		return x
	}

	if x, ok := v.(float32); ok {
		return float64(x)
	}

	if x, ok := v.(int); ok {
		return float64(x)
	}

	return 0
}

// SortAscBy Ordem crescente
func SortAscBy(d requester.StatusInvestResponse, name string) requester.StatusInvestResponse {
	sort.SliceStable(d, func(i, j int) bool {
		fpl := GetOrZero(reflections.GetField(&d[i], name))
		spl := GetOrZero(reflections.GetField(&d[j], name))

		return fpl < spl
	})

	return d
}

// SortDescBy Ordem decrescente
func SortDescBy(d requester.StatusInvestResponse, name string) requester.StatusInvestResponse {
	sort.SliceStable(d, func(i, j int) bool {
		fpl := GetOrZero(reflections.GetField(&d[i], name))
		spl := GetOrZero(reflections.GetField(&d[j], name))

		return fpl > spl
	})

	return d
}

// FindCompanyByName Busca empresa por nome
func FindCompanyByName(d requester.StatusInvestResponse, name string) int {
	for i := range d {
		if d[i].Ticker == name {
			return i
		}
	}

	return -1
}

// SortCompanies Ordena as companias pelos ranks
func SortCompanies(d requester.StatusInvestResponse) map[string]GreenblattRank {
	m := make(map[string]GreenblattRank)
	byPL := SortAscBy(d, "PL")
	SetRankNameBySlice(&m, byPL, "PlPosition")

	byRoic := SortDescBy(d, "Roic")
	SetRankNameBySlice(&m, byRoic, "RoicPosition")

	byDY := SortDescBy(d, "Dividend")
	SetRankNameBySlice(&m, byDY, "DyPosition")

	setSumRank(&m)
	return m
}

// SetRankNameBySlice Seta entradas do map pelo nome
func SetRankNameBySlice(m *map[string]GreenblattRank, d requester.StatusInvestResponse, fieldName string) {
	for _, v := range d {
		idx := FindCompanyByName(d, v.Ticker)
		createIfNotExists(m, v.Ticker)
		c := (*m)[v.Ticker]
		reflections.SetField(&c, fieldName, idx)
		(*m)[v.Ticker] = c
	}
}

func setSumRank(m *map[string]GreenblattRank) {
	for i, v := range *m {
		v.FinalPosition = v.DyPosition + v.PlPosition + v.RoicPosition
		(*m)[i] = v
	}
}

func createIfNotExists(m *map[string]GreenblattRank, name string) {
	if _, found := (*m)[name]; found == false {
		(*m)[name] = GreenblattRank{}
	}
}

func GetSortedByFinalPosition(m map[string]GreenblattRank, d requester.StatusInvestResponse) requester.StatusInvestResponse {
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

	s := requester.StatusInvestResponse{}

	for _, kv := range ss {
		idx := FindCompanyByName(d, kv.Key)
		c := d[idx]
		s = append(s, c)
	}

	return s
}
