package filterclasses

import (
	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/requester"
)

// GetFilters
func GetFilters(cfg config.ConfigSetup) []interface{ requester.AddFilterable } {
	var groupingClass []interface{ requester.AddFilterable }

	switch cfg.Grouping.Class {
	case "sector":
		groupingClass = GetSectorsAddFilterable()
	case "subsector":
		groupingClass = GetSubSectorsAddFilterable()
	case "segment":
		groupingClass = GetSegmentsAddFilterable()
	}

	return groupingClass
}
