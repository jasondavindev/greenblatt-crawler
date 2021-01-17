package filterclasses

import "github.com/jasondavindev/statusinvest-crawler/requester"

// Sector Sector filter class
type Sector struct {
	SectorID string
	Name     string
}

// ToAddFilter Parse sector to AdditionalFitler
func (s Sector) ToAddFilter() requester.AdditionalFilter {
	return requester.AdditionalFilter{
		FilterName:  "Sector",
		Value:       s.SectorID,
		Description: s.Name,
	}
}

// GetSectors Retorna todos os setores
func GetSectors() []Sector {
	return []Sector{
		{
			SectorID: "1",
			Name:     "Bens Industriais",
		},
		{
			SectorID: "2",
			Name:     "Consumo Cíclico",
		},
		{
			SectorID: "3",
			Name:     "Consumo não Cíclico",
		},
		{
			SectorID: "4",
			Name:     "Financeiro e Outros",
		},
		{
			SectorID: "5",
			Name:     "Materiais Básicos",
		},
		{
			SectorID: "6",
			Name:     "Petróleo. Gás e Biocombustíveis",
		},
		{
			SectorID: "7",
			Name:     "Saúde",
		},
		{
			SectorID: "8",
			Name:     "Tecnologia da Informação",
		},
		{
			SectorID: "9",
			Name:     "Comunicações",
		},
		{
			SectorID: "10",
			Name:     "Utilidade Pública",
		},
	}
}

// GetSectorsAddFilterable Return sectors as requester.AddFilterable
func GetSectorsAddFilterable() []interface{ requester.AddFilterable } {
	var filters []interface{ requester.AddFilterable }

	for _, s := range GetSectors() {
		filters = append(filters, s)
	}

	return filters
}
