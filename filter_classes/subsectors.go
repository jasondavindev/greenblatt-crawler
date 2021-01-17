package filterclasses

import "github.com/jasondavindev/statusinvest-crawler/requester"

// SubSector Subsetor
type SubSector struct {
	SubSectorID string
	SectorID    string
	Name        string
}

// ToAddFilter Parse sector to AdditionalFitler
func (s SubSector) ToAddFilter() requester.AdditionalFilter {
	return requester.AdditionalFilter{
		FilterName:  "Sector",
		Value:       s.SubSectorID,
		Description: s.Name,
	}
}

// GetSubSectorsAddFilterable Return SubSectors as requester.AddFilterable
func GetSubSectorsAddFilterable() []interface{ requester.AddFilterable } {
	var filters []interface{ requester.AddFilterable }

	for _, s := range GetSubSectors() {
		filters = append(filters, s)
	}

	return filters
}

// GetSubSectors Retorna todos os subsetores
func GetSubSectors() []SubSector {
	return []SubSector{
		{
			SubSectorID: "13",
			Name:        "Tecidos. Vestuário e Calçados",
		},
		{
			SubSectorID: "17",
			Name:        "Alimentos Processados",
		},
		{
			SubSectorID: "16",
			Name:        "Agropecuária",
		},
		{
			SubSectorID: "43",
			Name:        "Água e Saneamento",
		},
		{
			SubSectorID: "19",
			Name:        "Comércio e Distribuição",
		},
		{
			SubSectorID: "10",
			Name:        "Diversos",
		},
		{
			SubSectorID: "3",
			Name:        "Máquinas e Equipamentos",
		},
		{
			SubSectorID: "34",
			Name:        "Siderurgia e Metalurgia",
		},
		{
			SubSectorID: "15",
			Name:        "Viagens e Lazer",
		},
		{
			SubSectorID: "7",
			Name:        "Automóveis e Motocicletas",
		},
		{
			SubSectorID: "24",
			Name:        "Intermediários Financeiros",
		},
		{
			SubSectorID: "18",
			Name:        "Bebidas",
		},
		{
			SubSectorID: "59",
			Name:        "Biotecnologia",
		},
		{
			SubSectorID: "40",
			Name:        "Computadores e Equipamentos",
		},
		{
			SubSectorID: "2",
			Name:        "Construção e Engenharia",
		},
		{
			SubSectorID: "26",
			Name:        "Previdência e Seguros",
		},
		{
			SubSectorID: "8",
			Name:        "Comércio",
		},
		{
			SubSectorID: "14",
			Name:        "Utilidades Domésticas",
		},
		{
			SubSectorID: "29",
			Name:        "Embalagens",
		},
		{
			SubSectorID: "44",
			Name:        "Energia Elétrica",
		},
		{
			SubSectorID: "37",
			Name:        "Equipamentos",
		},
		{
			SubSectorID: "35",
			Name:        "Petróleo. Gás e Biocombustíveis",
		},
		{
			SubSectorID: "22",
			Name:        "Exploração de Imóveis",
		},
		{
			SubSectorID: "6",
			Name:        "Transporte",
		},
		{
			SubSectorID: "33",
			Name:        "Químicos",
		},
		{
			SubSectorID: "45",
			Name:        "Gás",
		},
		{
			SubSectorID: "28",
			Name:        "Serviços Financeiros Diversos",
		},
		{
			SubSectorID: "23",
			Name:        "Holdings Diversificadas",
		},
		{
			SubSectorID: "11",
			Name:        "Hoteis e Restaurantes",
		},
		{
			SubSectorID: "9",
			Name:        "Construção Civil",
		},
		{
			SubSectorID: "12",
			Name:        "Mídia",
		},
		{
			SubSectorID: "30",
			Name:        "Madeira e Papel",
		},
		{
			SubSectorID: "31",
			Name:        "Materiais Diversos",
		},
		{
			SubSectorID: "4",
			Name:        "Material de Transporte",
		},
		{
			SubSectorID: "1",
			Name:        "Comércio",
		},
		{
			SubSectorID: "36",
			Name:        "Comércio e Distribuição",
		},
		{
			SubSectorID: "38",
			Name:        "Medicamentos e Outros Produtos",
		},
		{
			SubSectorID: "32",
			Name:        "Mineração",
		},
		{
			SubSectorID: "25",
			Name:        "Outros",
		},
		{
			SubSectorID: "58",
			Name:        "Mídia",
		},
		{
			SubSectorID: "21",
			Name:        "Produtos de Uso Pessoal e de Limpeza",
		},
		{
			SubSectorID: "20",
			Name:        "Diversos",
		},
		{
			SubSectorID: "41",
			Name:        "Programas e Serviços",
		},
		{
			SubSectorID: "27",
			Name:        "Securitizadoras de Recebíveis",
		},
		{
			SubSectorID: "39",
			Name:        "Serv.Méd.Hospit..Análises e Diagnósticos",
		},
		{
			SubSectorID: "5",
			Name:        "Serviços",
		},
		{
			SubSectorID: "42",
			Name:        "Telecomunicações",
		},
	}
}
