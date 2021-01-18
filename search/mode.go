package search

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	filterclasses "github.com/jasondavindev/statusinvest-crawler/filter_classes"
	"github.com/jasondavindev/statusinvest-crawler/requester"
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
)

// Mode Search mode type
type Mode int

const (
	asyncSearchMode = 1
	syncSearchMode  = 0
)

// GetSearchMode Switch between async and sync search mode
func GetSearchMode(configSetup config.ConfigSetup) Mode {
	if !hasClassFilter(configSetup) && configSetup.Grouping.Class != "" {
		return asyncSearchMode
	}

	return syncSearchMode
}

// DoSearch Do request to api and return results
func DoSearch(searchMode Mode, configSetup config.ConfigSetup) []requester_types.StatusInvestResponseItem {
	if searchMode == asyncSearchMode {
		filters := filterclasses.GetFilters(configSetup)
		return requester.AsyncFind(&configSetup, filters)
	}

	return requester.SyncFind(&configSetup.Filters)
}

func hasClassFilter(configSetup config.ConfigSetup) bool {
	return configSetup.Filters["Segment"] != nil ||
		configSetup.Filters["Sector"] != nil ||
		configSetup.Filters["SubSector"] != nil
}
