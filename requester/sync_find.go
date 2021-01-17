package requester

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
)

// SyncFind Make Sync requests to API
func SyncFind(cfg *config.Filters) StatusInvestResponse {
	reqParams := ParseFiltersToParams(cfg)
	json := ToJSON(reqParams)

	url := CreateURL(*cfg, json)
	res := DoRequest(url)

	defer res.Body.Close()

	return ParseResponse(res.Body)
}
