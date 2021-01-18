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
	FilterName  string
	Value       string
	Description string
}

// AddFilterable Filterable classes
type AddFilterable interface {
	ToAddFilter() AdditionalFilter
}

// AsyncFind Make parallel requests to API and return results
func AsyncFind(configSetup *config.ConfigSetup, addFilterClass []interface{ AddFilterable }) requester_types.StatusInvestResponse {
	var wg sync.WaitGroup
	var ch chan requester_types.StatusInvestResponse

	filters := setupAddFilters(addFilterClass)
	length := len(filters)

	ch = make(chan requester_types.StatusInvestResponse, length)
	wg.Add(length)

	for _, filter := range filters {
		go doAsyncFind(&wg, &ch, configSetup, filter)
	}

	wg.Wait()
	close(ch)

	return readResults(ch, &configSetup.Grouping)
}

func doAsyncFind(wg *sync.WaitGroup, ch *chan requester_types.StatusInvestResponse, configSetup *config.ConfigSetup, addFilter AdditionalFilter) {
	defer (*wg).Done()

	filters := copyMapFiltersAndAddClassFilter(&configSetup.Filters, &addFilter)
	response := SyncFind(&filters)
	results := addClassToResults(response, addFilter)
	*ch <- results
	fmt.Printf("[%s] encontrado %d resultados\n", addFilter.Description, len(response))
}

func addClassToResults(results requester_types.StatusInvestResponse, addFilter AdditionalFilter) requester_types.StatusInvestResponse {
	for k := range results {
		results[k].Area = addFilter.Description
	}

	return results
}

func copyMapFiltersAndAddClassFilter(configSetup *config.Filters, addFilter *AdditionalFilter) config.Filters {
	mapCopy := config.Filters{}

	for k, v := range *configSetup {
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
	countTopResults := getCountTopResults(grouping.Top)

	for response := range ch {
		uniqueCompanies := removeDuplicateCompanies(response)
		sortedCompanies := greenblatt.SortCompanies(uniqueCompanies)
		finalPosition := greenblatt.GetSortedByFinalPosition(sortedCompanies, response)
		topResults := getTopResults(finalPosition, countTopResults)
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

func getCountTopResults(top int) int {
	if top <= 0 {
		return 5
	}

	return top
}

func removeDuplicateCompanies(companies requester_types.StatusInvestResponse) requester_types.StatusInvestResponse {
	results := requester_types.StatusInvestResponse{}
	resultsMap := map[string]requester_types.StatusInvestResponseItem{}
	const TickerONType = "3"

	for _, company := range companies {
		companyTickPrefix := company.Ticker[0:4]
		tickerType := company.Ticker[4:]

		if _, ok := resultsMap[companyTickPrefix]; !ok ||
			(ok && (tickerType == TickerONType ||
				company.PL < resultsMap[companyTickPrefix].PL ||
				company.Preco < resultsMap[companyTickPrefix].Preco)) {
			resultsMap[companyTickPrefix] = company
		}
	}

	for _, company := range resultsMap {
		results = append(results, company)
	}

	return results
}
