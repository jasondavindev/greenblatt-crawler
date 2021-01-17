package requester

import (
	"sync"

	"github.com/jasondavindev/statusinvest-crawler/config"
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
func AsyncFind(cfg *config.Filters, addFilterClass []interface{ AddFilterable }) StatusInvestResponse {
	var wg sync.WaitGroup
	var ch chan StatusInvestResponse

	filters := setupAddFilters(addFilterClass)
	length := len(filters)

	ch = make(chan StatusInvestResponse, length)
	wg.Add(length)

	for _, filter := range filters {
		go doAsyncFind(&wg, &ch, cfg, filter)
	}

	wg.Wait()
	close(ch)

	return readResults(ch)
}

func doAsyncFind(wg *sync.WaitGroup, ch *chan StatusInvestResponse, cfg *config.Filters, addFilter AdditionalFilter) {
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

func readResults(ch chan StatusInvestResponse) StatusInvestResponse {
	results := StatusInvestResponse{}

	for response := range ch {
		results = append(results, response...)
	}

	return results
}
