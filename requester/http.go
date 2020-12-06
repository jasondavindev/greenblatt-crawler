package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

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
}

// StatusInvestResponse Resposta status invest
type StatusInvestResponse = []StatusInvestResponseItem

// Find Faz solicitacao para site status invest
func Find(r RequestParams) StatusInvestResponse {
	fmt.Println("Buscando recursos na API Status Invest...")

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

// ToJSON Converte request params para formato JSON
func (crp *RequestParams) ToJSON() string {
	jsonMap, err := json.Marshal(crp)

	if err != nil {
		panic(err)
	}

	return string(jsonMap)
}

// CreateURL Cria string content URL de request
func (crp *RequestParams) CreateURL() string {
	return fmt.Sprintf("https://statusinvest.com.br/category/advancedsearchresult?CategoryType=1&search=%s", url.QueryEscape(crp.ToJSON()))
}
