package search

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	filterclasses "github.com/jasondavindev/statusinvest-crawler/filter_classes"
	"github.com/jasondavindev/statusinvest-crawler/requester"
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
)

// GetSearchMode Switch between async and sync search mode
func GetSearchMode(configSetup config.ConfigSetup) []requester_types.StatusInvestResponseItem {
	if configSetup.Grouping.Class != "" {
		filters := filterclasses.GetFilters(configSetup)
		return requester.AsyncFind(&configSetup, filters)
	}

	return requester.SyncFind(&configSetup.Filters)
}
