package requester

import (
	"fmt"
	"sync"

	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/greenblatt"
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
)

// AdditionalFilter Filter type to async find
type AdditionalFilter struct {
	FilterName string
	Value      string
}

// AddFilterable Filterable classes
type AddFilterable interface {
	ToAddFilter() AdditionalFilter
}

// AsyncFind Make parallel requests to API and return results
func AsyncFind(cfg *config.ConfigSetup, addFilterClass []interface{ AddFilterable }) requester_types.StatusInvestResponse {
	var wg sync.WaitGroup
	var ch chan requester_types.StatusInvestResponse

	filters := setupAddFilters(addFilterClass)
	length := len(filters)

	ch = make(chan requester_types.StatusInvestResponse, length)
	wg.Add(length)

	for _, filter := range filters {
		go doAsyncFind(&wg, &ch, &cfg.Filters, filter)
	}

	wg.Wait()
	close(ch)

	return readResults(ch, &cfg.Grouping)
}

func doAsyncFind(wg *sync.WaitGroup, ch *chan requester_types.StatusInvestResponse, cfg *config.Filters, addFilter AdditionalFilter) {
	defer (*wg).Done()

	filters := copyMapFiltersAndAddClassFilter(cfg, &addFilter)
	response := SyncFind(&filters)
	*ch <- response
}

func copyMapFiltersAndAddClassFilter(cfg *config.Filters, addFilter *AdditionalFilter) config.Filters {
	mapCopy := config.Filters{}

	for k, v := range *cfg {
		mapCopy[k] = v
	}

	mapCopy[(*addFilter).FilterName] = (*addFilter).Value

	return mapCopy
}

func setupAddFilters(addFilterClass []interface{ AddFilterable }) []AdditionalFilter {
	filters := make([]AdditionalFilter, len(addFilterClass))

	for i, filter := range addFilterClass {
		filters[i] = filter.ToAddFilter()
	}

	return filters
}

func readResults(ch chan requester_types.StatusInvestResponse, grouping *config.Grouping) requester_types.StatusInvestResponse {
	results := requester_types.StatusInvestResponse{}

	for response := range ch {
		m := greenblatt.SortCompanies(response)
		fp := greenblatt.GetSortedByFinalPosition(m, response)
		topResults := getTopResults(fp, grouping.Top)
		fmt.Println("Top: ", len(topResults))
		results = append(results, topResults...)
	}

	return results
}

func getTopResults(results requester_types.StatusInvestResponse, top int) requester_types.StatusInvestResponse {
	topResults := requester_types.StatusInvestResponse{}
	i := 0

	for _, v := range results {
		if i >= top {
			break
		}

		topResults = append(topResults, v)
		i++
	}

	return topResults
}
