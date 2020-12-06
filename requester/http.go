package requester

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// StatusInvestResponseItem Item resposta
type StatusInvestResponseItem struct {
	CompanyName                    string  `json:"companyName"`
	Ticker                         string  `json:"ticker"`
	Price                          float32 `json:"price"`
	PL                             float32 `json:"p_L"`
	Dividend                       float32 `json:"dy"`
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
}

// StatusInvestResponse Resposta status invest
type StatusInvestResponse = []StatusInvestResponseItem

// Find Faz solicitacao para site status invest
func Find(r RequestParams) StatusInvestResponse {
	getURL := r.CreateURL()

	res, err := http.Get(getURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	return ParseResponse(res.Body)
}

// ParseResponse Converte resposta string para objeto
func ParseResponse(res io.Reader) StatusInvestResponse {
	body, err := ioutil.ReadAll(res)

	if err != nil {
		log.Fatalln(err)
	}

	var b StatusInvestResponse

	if err := json.Unmarshal(body, &b); err != nil {
		log.Fatalln(err)
	}

	return b
}
