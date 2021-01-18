package requester_types

// StatusInvestResponseItem Item resposta
type StatusInvestResponseItem struct {
	Companhia                      string  `json:"companyName"`
	Ticker                         string  `json:"ticker"`
	Preco                          float32 `json:"price"`
	PL                             float32 `json:"p_L"`
	Dividendo                      float32 `json:"dy"`
	Roic                           float32 `json:"roic"`
	DividaLiquidaEbit              float32 `json:"dividaLiquidaEbit"`
	DividaliquidaPatrimonioLiquido float32 `json:"dividaliquidaPatrimonioLiquido"`
	EVEbit                         float32 `json:"eV_Ebit"`
	LiquidezMediaDiaria            float32 `json:"liquidezMediaDiaria"`
	Lpa                            float32 `json:"lpa"`
	LucrosCagr5                    float32 `json:"lucros_Cagr5"`
	MargemBruta                    float32 `json:"margemBruta"`
	MargemEbit                     float32 `json:"margemEbit"`
	MargemLiquida                  float32 `json:"margemLiquida"`
	PEbit                          float32 `json:"p_Ebit"`
	PVP                            float32 `json:"p_VP"`
	ReceitasCagr5                  float32 `json:"receitas_Cagr5"`
	Roa                            float32 `json:"roa"`
	Roe                            float32 `json:"roe"`
	Vpa                            float32 `json:"vpa"`
	Area                           string
}

// StatusInvestResponse Resposta status invest
type StatusInvestResponse = []StatusInvestResponseItem
