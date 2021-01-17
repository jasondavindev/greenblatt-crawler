package requester

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
)

// SyncFind Make Sync requests to API
func SyncFind(cfg *config.Filters) requester_types.StatusInvestResponse {
	reqParams := ParseFiltersToParams(cfg)
	json := ToJSON(reqParams)

	url := CreateURL(*cfg, json)
	res := DoRequest(url)

	defer res.Body.Close()

	return ParseResponse(res.Body)
}
