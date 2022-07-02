package marketdata

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	config "github.com/PascalSchroederDE/go-market-lemonizer/configs"
	api_helper "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/api"
	datastructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
	lemon_interval "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs/interval"
)

// Constructor

type LemonMarketAPI struct {
	apiKey  string
	baseUrl string
}

func Init(apiKey string) LemonMarketAPI {
	return LemonMarketAPI{apiKey: apiKey, baseUrl: config.LEMON_MARKET_DATA_ENDPOINT}
}

func InitManualEndpoint(apiKey string, endpoint string) LemonMarketAPI {
	return LemonMarketAPI{apiKey: apiKey, baseUrl: endpoint}
}

// Exposed API connector functions

func (impl *LemonMarketAPI) getData(isin []string, interval lemon_interval.LemonInterval, parameter map[string]string) (map[string][]datastructs.OHLCData, error) {
	var isinSets [][]string = buildCompatibleIsinSets(isin)
	var stockData map[string][]datastructs.OHLCData = make(map[string][]datastructs.OHLCData)

	for _, isinSet := range isinSets {
		var page int = 1
		for {
			parameter["isin"] = strings.Join(isinSet, ",")
			parameter["page"] = fmt.Sprint(page)

			req, err := api_helper.BuildGetRequestWithParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.OHLC_ENDPOINT+interval.String()+"/", parameter)
			if err != nil {
				return nil, errors.New("build request error")
			}
			respObj, err := performStockDataRequest(req)
			if err != nil || reflect.DeepEqual(respObj, getOhlcDataResponse{}) {
				return nil, errors.New("bad response error")
			}
			mapResultsToIsinMap(stockData, respObj.Results)
			if respObj.Pages <= respObj.Page {
				break
			}

			page++
		}
	}

	return stockData, nil
}

func (impl *LemonMarketAPI) GetLatestData(isin []string, interval lemon_interval.LemonInterval) (map[string][]datastructs.OHLCData, error) {
	parameter := map[string]string{
		"limit": "100",
	}

	return impl.getData(isin, interval, parameter)
}

func (impl *LemonMarketAPI) GetHistoryData(isin []string, dateFrom time.Time, dateTo time.Time, interval lemon_interval.LemonInterval) (map[string][]datastructs.OHLCData, error) {
	if interval == lemon_interval.Min || interval == lemon_interval.Hour {
		yearFrom, monthFrom, dayFrom := dateFrom.Date()
		yearTo, monthTo, dayTo := dateTo.Date()
		if yearFrom != yearTo || monthFrom != monthTo || dayFrom != dayTo {
			return impl.executeRequestsPerDay(dateFrom, dateTo, isin, interval)
		}
	}

	parameter := map[string]string{
		"limit": "100",
		"from":  dateFrom.Format("2006-01-02T15:04:05"),
		"to":    dateTo.Format("2006-01-02T15:04:05"),
	}

	return impl.getData(isin, interval, parameter)
}

func (impl *LemonMarketAPI) GetVenueInformation(mic string) (datastructs.VenueResult, error) {
	parameter := map[string]string{
		"limit": "100",
		"mic":   mic,
	}

	req, err := api_helper.BuildGetRequestWithParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.VENUES_ENDPOINT, parameter)
	if err != nil {
		return datastructs.VenueResult{}, errors.New("build request error")
	}
	respObj, err := performVenueDataRequest(req)
	if err != nil || reflect.DeepEqual(respObj, venueResponse{}) {
		return datastructs.VenueResult{}, errors.New("bad response error")
	}

	for _, result := range respObj.Results {
		if strings.EqualFold(result.Mic, mic) {
			return result, nil
		}
	}
	return datastructs.VenueResult{}, errors.New("mic not found")
}
