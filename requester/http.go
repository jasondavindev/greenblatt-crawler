package requester

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
)

// Find Faz solicitacao para site status invest
func Find(cfg config.Filters) requester_types.StatusInvestResponse {
	fmt.Println("Buscando recursos na API Status Invest...")

	reqParams := ParseFiltersToParams(&cfg)
	json := ToJSON(reqParams)

	url := CreateURL(cfg, json)
	res := DoRequest(url)

	defer res.Body.Close()

	return ParseResponse(res.Body)
}

// DoRequest Make de request to API
func DoRequest(url string) *http.Response {
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	return res
}

// ParseResponse Converte resposta string para objeto
func ParseResponse(res io.Reader) requester_types.StatusInvestResponse {
	body, err := ioutil.ReadAll(res)

	if err != nil {
		log.Fatalln(err)
	}

	var b requester_types.StatusInvestResponse

	if err := json.Unmarshal(body, &b); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Encontrado %d resultados\n", len(b))

	return b
}

// CreateURL Cria string content URL de request
func CreateURL(cfg config.Filters, json string) string {
	return fmt.Sprintf("https://statusinvest.com.br/category/advancedsearchresult?CategoryType=1&search=%s", url.QueryEscape(json))
}

// ToJSON Converte request params para formato JSON
func ToJSON(v interface{}) string {
	jsonMap, err := json.Marshal(v)

	if err != nil {
		panic(err)
	}

	return string(jsonMap)
}
