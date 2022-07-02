package marketdata

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	api_helper "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/api"
	datastructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
	lemon_interval "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs/interval"
)

// Define API JSON response to Go Object

type getOhlcDataResponse struct {
	Results  []ohlcDataResult `json:"results"`
	Previous interface{}      `json:"previous"`
	Next     interface{}      `json:"next"`
	Total    int              `json:"total"`
	Page     int              `json:"page"`
	Pages    int              `json:"pages"`
}

type ohlcDataResult struct {
	Isin string    `json:"isin"`
	O    float64   `json:"o"`
	H    float64   `json:"h"`
	L    float64   `json:"l"`
	C    float64   `json:"c"`
	T    time.Time `json:"t"`
	Mic  string    `json:"mic"`
}

type venueResponse struct {
	Results  []datastructs.VenueResult `json:"results"`
	Previous interface{}               `json:"previous"`
	Next     interface{}               `json:"next"`
	Total    int                       `json:"total"`
	Page     int                       `json:"page"`
	Pages    int                       `json:"pages"`
}

// helper functions

func performStockDataRequest(req *http.Request) (*getOhlcDataResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj getOhlcDataResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}
	return &respObj, nil
}

func performVenueDataRequest(req *http.Request) (*venueResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj venueResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}

	return &respObj, nil
}

// helper for summarizing complex requests with lots of received data

func (impl *LemonMarketAPI) executeRequestsPerDay(dateFrom time.Time, dateTo time.Time, isin []string, interval lemon_interval.LemonInterval) (map[string][]datastructs.OHLCData, error) {
	yearTo, monthTo, dayTo := dateTo.Date()
	dateToStep := time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 23, 59, 59, 0, dateFrom.Location())
	stop := false
	mapCollection := make(map[string][]datastructs.OHLCData)
	for {
		yearStep, monthStep, dayStep := dateToStep.Date()
		if yearStep == yearTo && monthStep == monthTo && dayStep == dayTo {
			dateToStep = dateTo
			stop = true
		}

		parameter := map[string]string{
			"limit": "100",
			"from":  dateFrom.Format("2006-01-02T15:04:05"),
			"to":    dateToStep.Format("2006-01-02T15:04:05"),
		}

		newMap, err := impl.getData(isin, interval, parameter)
		if err != nil {
			return nil, err
		}

		mergeIsinMaps(mapCollection, newMap)

		dateFrom = time.Date(dateToStep.Year(), dateToStep.Month(), dateToStep.Day()+1, 0, 0, 1, dateToStep.Nanosecond(), dateToStep.Location())
		dateToStep = time.Date(dateToStep.Year(), dateToStep.Month(), dateToStep.Day()+1, dateToStep.Hour(), dateToStep.Minute(), dateToStep.Second(), dateToStep.Nanosecond(), dateToStep.Location())
		if stop {
			break
		}
	}
	return mapCollection, nil
}

// market api limited util functions

func buildCompatibleIsinSets(isins []string) [][]string {
	var isinSets [][]string
	for i, isin := range isins {
		if i%10 == 0 || i == 0 {
			isinSets = append(isinSets, []string{isin})
		} else {
			isinSets[i/10] = append(isinSets[i/10], isin)
		}
	}
	return isinSets
}

func mapResultsToIsinMap(existingMap map[string][]datastructs.OHLCData, newSet []ohlcDataResult) {
	for _, datapoint := range newSet {
		if val, ok := existingMap[datapoint.Isin]; ok {
			existingMap[datapoint.Isin] = append(val, datastructs.OHLCData{Time: datapoint.T, High: datapoint.H, Low: datapoint.L, Opening: datapoint.O, Closing: datapoint.C})
		} else {
			existingMap[datapoint.Isin] = []datastructs.OHLCData{{Time: datapoint.T, High: datapoint.H, Low: datapoint.L, Opening: datapoint.O, Closing: datapoint.C}}
		}
	}
}

func mergeIsinMaps(map1 map[string][]datastructs.OHLCData, map2 map[string][]datastructs.OHLCData) {
	for isin, datapoints := range map2 {
		if val, ok := map1[isin]; ok {
			map1[isin] = append(val, datapoints...)
		} else {
			map1[isin] = datapoints
		}
	}
}
