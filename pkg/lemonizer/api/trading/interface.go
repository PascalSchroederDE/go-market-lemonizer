package trading

import (
	"errors"
	"fmt"
	"time"

	config "github.com/PascalSchroederDE/go-market-lemonizer/configs"
	api_helper "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/api"
	utils "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/utils"
	datastructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
)

// Constructor

type LemonTradingAPI struct {
	apiKey  string
	baseUrl string
}

func Init(apiKey string) LemonTradingAPI {
	return LemonTradingAPI{apiKey: apiKey, baseUrl: config.LEMON_MARKET_TRADING_ENDPOINT}
}

func InitPaperEndpoint(apiKey string) LemonTradingAPI {
	return LemonTradingAPI{apiKey: apiKey, baseUrl: config.LEMON_MARKET_PAPER_TRADING_ENDPOINT}
}

func InitManualEndpoint(apiKey string, endpoint string) LemonTradingAPI {
	return LemonTradingAPI{apiKey: apiKey, baseUrl: endpoint}
}

// Exposed API connector functions

func (impl LemonTradingAPI) placeStockTitleOrder(isin string, amount int, limitPrice int, expiryTime time.Time, mode string) (datastructs.PlaceOrderResults, error) {
	parameter := map[string]string{
		"expires_at":  expiryTime.Format("2006-01-02T15:04:05"),
		"side":        mode,
		"quantity":    fmt.Sprint(amount),
		"venue":       "xmun",
		"isin":        isin,
		"limit_price": fmt.Sprint(limitPrice),
	}

	req, err := api_helper.BuildPostRequestWithParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.ORDERS_ENDPOINT, parameter)
	if err != nil {
		return datastructs.PlaceOrderResults{}, errors.New("build request error")
	}
	respObj, err := performPlaceOrderRequest(req)
	if err != nil {
		return datastructs.PlaceOrderResults{}, errors.New("bad response error")
	}

	if respObj.Status == "ok" {
		return respObj.Results, nil
	}
	return datastructs.PlaceOrderResults{}, errors.New("order not ok - status: " + respObj.Status + "\nResponse Object: " + utils.PrettyPrint(respObj))
}

func (impl LemonTradingAPI) PlaceSellStockTitleOrder(isin string, amount int, limitPrice int, expiryTime time.Time) (datastructs.PlaceOrderResults, error) {
	return impl.placeStockTitleOrder(isin, amount, limitPrice, expiryTime, "sell")
}

func (impl LemonTradingAPI) PlaceBuyStockTitleOrder(isin string, amount int, limitPrice int, expiryTime time.Time) (datastructs.PlaceOrderResults, error) {
	return impl.placeStockTitleOrder(isin, amount, limitPrice, expiryTime, "buy")
}

func (impl LemonTradingAPI) ActivateOrder(orderId string) error {
	req, err := api_helper.BuildPostRequestWithoutParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.ORDERS_ENDPOINT+orderId+"/activate/")
	if err != nil {
		return errors.New("build request error")
	}
	respObj, err := performActivateOrderRequest(req)
	if err != nil {
		return errors.New("bad response error")
	}

	if respObj.Status == "ok" {
		return nil
	}
	return errors.New("activation not ok")
}

func (impl LemonTradingAPI) DeleteOrder(orderId string) error {
	req, err := api_helper.BuildDeleteRequestWithoutParameter(impl.apiKey, impl.baseUrl, config.VERSION+config.ORDERS_ENDPOINT+orderId)
	if err != nil {
		return errors.New("build request error")
	}
	respObj, err := performDeleteOrderRequest(req)
	if err != nil {
		return errors.New("bad response error")
	}

	if respObj.Status == "ok" {
		return nil
	}
	return errors.New("deletion not ok")
}
