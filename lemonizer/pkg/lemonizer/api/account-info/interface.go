package accountinfo

import (
	"errors"
	"fmt"

	config "github.com/PascalSchroederDE/go-market-lemonizer/configs"
	api_helper "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/api"
	datastructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
)

// Constructor

type lemonAccountAPI struct {
	apiKey  string
	baseUrl string
}

func Init(apiKey string) lemonAccountAPI {
	return lemonAccountAPI{apiKey: apiKey, baseUrl: config.LEMON_MARKET_TRADING_ENDPOINT}
}

func InitPaperEndpoint(apiKey string) lemonAccountAPI {
	return lemonAccountAPI{apiKey: apiKey, baseUrl: config.LEMON_MARKET_PAPER_TRADING_ENDPOINT}
}

func InitManualEndpoint(apiKey string, endpoint string) lemonAccountAPI {
	return lemonAccountAPI{apiKey: apiKey, baseUrl: endpoint}
}

// Exposed API connector functions

func (impl lemonAccountAPI) GetPortfolio() (map[string]datastructs.PortfolioResults, error) {
	var portfolio map[string]datastructs.PortfolioResults = make(map[string]datastructs.PortfolioResults)

	var page int = 1
	for {
		parameter := map[string]string{
			"page": fmt.Sprint(page),
		}

		req, err := api_helper.BuildGetRequestWithParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.POSITIONS_ENDPOINT, parameter)
		if err != nil {
			return nil, errors.New("build request error")
		}
		respObj, err := performPortfolioRequest(req)
		if err != nil || respObj.Status != "ok" {
			return nil, errors.New("bad response error")
		}

		for _, result := range respObj.Results {
			portfolio[result.Isin] = result
		}

		if respObj.Pages <= respObj.Page {
			break
		}
		page++
	}

	return portfolio, nil
}

func (impl lemonAccountAPI) GetAccount() (datastructs.GetAccountResult, error) {
	req, err := api_helper.BuildGetRequestWithoutParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.ACCOUNT_ENDPOINT)
	if err != nil {
		return datastructs.GetAccountResult{}, errors.New("build request error")
	}
	respObj, err := performAccountRequest(req)
	if err != nil {
		return datastructs.GetAccountResult{}, errors.New("bad response error")
	}
	return respObj.Results, nil
}
