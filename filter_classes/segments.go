package filterclasses

import "github.com/jasondavindev/statusinvest-crawler/requester"

// Segment Segmento
type Segment struct {
	SegmentID string
	Name      string
}

// ToAddFilter Parse sector to AdditionalFitler
func (s Segment) ToAddFilter() requester.AdditionalFilter {
	return requester.AdditionalFilter{
		FilterName:  "Sector",
		Value:       s.SegmentID,
		Description: s.Name,
	}
}

// GetSegmentsAddFilterable Return Segments as requester.AddFilterable
func GetSegmentsAddFilterable() []interface{ requester.AddFilterable } {
	var filters []interface{ requester.AddFilterable }

	for _, s := range GetSegments() {
		filters = append(filters, s)
	}

	return filters
}

// GetSegments Retorna todos os segmentos
func GetSegments() []Segment {
	return []Segment{
		{
			SegmentID: "31",
			Name:      "Acessórios",
		},
		{
			SegmentID: "43",
			Name:      "Açucar e Alcool",
		},
		{
			SegmentID: "42",
			Name:      "Agricultura",
		},
		{
			SegmentID: "84",
			Name:      "Água e Saneamento",
		},
		{
			SegmentID: "47",
			Name:      "Alimentos",
		},
		{
			SegmentID: "44",
			Name:      "Alimentos Diversos",
		},
		{
			SegmentID: "24",
			Name:      "Aluguel de carros",
		},
		{
			SegmentID: "6",
			Name:      "Armas e Munições",
		},
		{
			SegmentID: "72",
			Name:      "Artefatos de Cobre",
		},
		{
			SegmentID: "73",
			Name:      "Artefatos de Ferro e Aço",
		},
		{
			SegmentID: "101",
			Name:      "Atividades Esportivas",
		},
		{
			SegmentID: "19",
			Name:      "Automóveis e Motocicletas",
		},
		{
			SegmentID: "54",
			Name:      "Bancos",
		},
		{
			SegmentID: "104",
			Name:      "Bebida Alcoólicas",
		},
		{
			SegmentID: "38",
			Name:      "Bicicletas",
		},
		{
			SegmentID: "106",
			Name:      "Biotecnologia",
		},
		{
			SegmentID: "39",
			Name:      "Brinquedos e Jogos",
		},
		{
			SegmentID: "32",
			Name:      "Calçados",
		},
		{
			SegmentID: "45",
			Name:      "Carnes e Derivados",
		},
		{
			SegmentID: "46",
			Name:      "Cervejas e Refrigerantes",
		},
		{
			SegmentID: "81",
			Name:      "Computadores e Equipamentos",
		},
		{
			SegmentID: "2",
			Name:      "Construção Pesada",
		},
		{
			SegmentID: "58",
			Name:      "Corretoras de Seguros",
		},
		{
			SegmentID: "20",
			Name:      "Eletrodomésticos",
		},
		{
			SegmentID: "35",
			Name:      "Eletrodomésticos",
		},
		{
			SegmentID: "63",
			Name:      "Embalagens",
		},
		{
			SegmentID: "85",
			Name:      "Energia Elétrica",
		},
		{
			SegmentID: "3",
			Name:      "Engenharia Consultiva",
		},
		{
			SegmentID: "78",
			Name:      "Equipamentos",
		},
		{
			SegmentID: "75",
			Name:      "Equipamentos e Serviços",
		},
		{
			SegmentID: "51",
			Name:      "Exploração de Imóveis",
		},
		{
			SegmentID: "13",
			Name:      "Exploração de Rodovias",
		},
		{
			SegmentID: "76",
			Name:      "Exploração. Refino e Distribuição",
		},
		{
			SegmentID: "69",
			Name:      "Fertilizantes e Defensivos",
		},
		{
			SegmentID: "33",
			Name:      "Fios e Tecidos",
		},
		{
			SegmentID: "86",
			Name:      "Gás",
		},
		{
			SegmentID: "61",
			Name:      "Gestão de Recursos e Investimentos",
		},
		{
			SegmentID: "53",
			Name:      "Holdings Diversificadas",
		},
		{
			SegmentID: "27",
			Name:      "Hotelaria",
		},
		{
			SegmentID: "23",
			Name:      "Incorporações",
		},
		{
			SegmentID: "52",
			Name:      "Intermediação Imobiliária",
		},
		{
			SegmentID: "29",
			Name:      "Jornais. Livros e Revistas",
		},
		{
			SegmentID: "108",
			Name:      "Logística",
		},
		{
			SegmentID: "64",
			Name:      "Madeira",
		},
		{
			SegmentID: "7",
			Name:      "Máq. e Equip. Construção e Agrícolas",
		},
		{
			SegmentID: "8",
			Name:      "Máq. e Equip. Industriais",
		},
		{
			SegmentID: "66",
			Name:      "Materiais Diversos",
		},
		{
			SegmentID: "10",
			Name:      "Material Aeronáutico e de Defesa",
		},
		{
			SegmentID: "1",
			Name:      "Material de Transporte",
		},
		{
			SegmentID: "11",
			Name:      "Material Rodoviário",
		},
		{
			SegmentID: "77",
			Name:      "Medicamentos e Outros Produtos",
		},
		{
			SegmentID: "79",
			Name:      "Medicamentos e Outros Produtos",
		},
		{
			SegmentID: "67",
			Name:      "Minerais Metálicos",
		},
		{
			SegmentID: "68",
			Name:      "Minerais Não Metálicos",
		},
		{
			SegmentID: "9",
			Name:      "Motores . Compressores e Outros",
		},
		{
			SegmentID: "36",
			Name:      "Móveis",
		},
		{
			SegmentID: "57",
			Name:      "Outros",
		},
		{
			SegmentID: "65",
			Name:      "Papel e Celulose",
		},
		{
			SegmentID: "70",
			Name:      "Petroquímicos",
		},
		{
			SegmentID: "40",
			Name:      "Produção de Eventos e Shows",
		},
		{
			SegmentID: "30",
			Name:      "Produção e Difusão de Filmes e Programas",
		},
		{
			SegmentID: "105",
			Name:      "Produção e Difusão de Filmes e Programas",
		},
		{
			SegmentID: "49",
			Name:      "Produtos de Limpeza",
		},
		{
			SegmentID: "50",
			Name:      "Produtos de Uso Pessoal",
		},
		{
			SegmentID: "21",
			Name:      "Produtos Diversos",
		},
		{
			SegmentID: "48",
			Name:      "Produtos Diversos",
		},
		{
			SegmentID: "4",
			Name:      "Produtos para Construção",
		},
		{
			SegmentID: "25",
			Name:      "Programas de Fidelização",
		},
		{
			SegmentID: "82",
			Name:      "Programas e Serviços",
		},
		{
			SegmentID: "107",
			Name:      "Publicidade",
		},
		{
			SegmentID: "71",
			Name:      "Químicos Diversos",
		},
		{
			SegmentID: "28",
			Name:      "Restaurante e Similares",
		},
		{
			SegmentID: "60",
			Name:      "Securitizadoras de Recebíveis",
		},
		{
			SegmentID: "59",
			Name:      "Seguradoras",
		},
		{
			SegmentID: "80",
			Name:      "Serv.Méd.Hospit..Análises e Diagnósticos",
		},
		{
			SegmentID: "14",
			Name:      "Serviços de Apoio e Armazenagem",
		},
		{
			SegmentID: "5",
			Name:      "Serviços Diversos",
		},
		{
			SegmentID: "12",
			Name:      "Serviços Diversos",
		},
		{
			SegmentID: "26",
			Name:      "Serviços Educacionais",
		},
		{
			SegmentID: "62",
			Name:      "Serviços Financeiros Diversos",
		},
		{
			SegmentID: "74",
			Name:      "Siderurgia",
		},
		{
			SegmentID: "55",
			Name:      "Soc. Arrendamento Mercantil",
		},
		{
			SegmentID: "56",
			Name:      "Soc. Crédito e Financiamento",
		},
		{
			SegmentID: "102",
			Name:      "Tabaco",
		},
		{
			SegmentID: "22",
			Name:      "Tecidos. Vestuário e Calçados",
		},
		{
			SegmentID: "83",
			Name:      "Telecomunicações",
		},
		{
			SegmentID: "15",
			Name:      "Transporte Aéreo",
		},
		{
			SegmentID: "16",
			Name:      "Transporte Ferroviário",
		},
		{
			SegmentID: "17",
			Name:      "Transporte Hidroviário",
		},
		{
			SegmentID: "18",
			Name:      "Transporte Rodoviário",
		},
		{
			SegmentID: "37",
			Name:      "Utensílios Domésticos",
		},
		{
			SegmentID: "34",
			Name:      "Vestuário",
		},
		{
			SegmentID: "41",
			Name:      "Viagens e Turismo",
		},
	}
}
