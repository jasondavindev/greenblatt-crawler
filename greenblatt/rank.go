package greenblatt

import (
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
	"github.com/oleiade/reflections"
)

// RankItem Estrutura do item no rank
type RankItem struct {
	PlPosition    int
	RoicPosition  int
	DyPosition    int
	FinalPosition int
}

// SetRankNameBySlice Seta entradas do map pelo nome
func SetRankNameBySlice(m *map[string]RankItem, d requester_types.StatusInvestResponse, fieldName string) {
	for _, v := range d {
		idx := FindCompanyByName(d, v.Ticker)
		CreateKeyIfNotExists(m, v.Ticker)
		c := (*m)[v.Ticker]
		reflections.SetField(&c, fieldName, idx)
		(*m)[v.Ticker] = c
	}
}

// CalculateFinalRankField Soma as posicoes de cada item para setar a posição final
func CalculateFinalRankField(m *map[string]RankItem) {
	for i, v := range *m {
		v.FinalPosition = v.DyPosition + v.PlPosition + v.RoicPosition
		(*m)[i] = v
	}
}

// CreateKeyIfNotExists Cria chave se ticker nao existe
func CreateKeyIfNotExists(m *map[string]RankItem, name string) {
	if _, found := (*m)[name]; found == false {
		(*m)[name] = RankItem{}
	}
}
