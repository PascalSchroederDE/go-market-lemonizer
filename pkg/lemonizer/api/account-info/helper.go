package accountinfo

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	api_helper "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/api"
	datastructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
)

// Define API JSON response to Go Object

type getPortfolioResponse struct {
	Time     time.Time                      `json:"time"`
	Status   string                         `json:"status"`
	Mode     string                         `json:"mode"`
	Results  []datastructs.PortfolioResults `json:"results"`
	Previous interface{}                    `json:"previous"`
	Next     interface{}                    `json:"next"`
	Total    int                            `json:"total"`
	Page     int                            `json:"page"`
	Pages    int                            `json:"pages"`
}

type getAccountResponse struct {
	Time    time.Time                    `json:"time"`
	Mode    string                       `json:"mode"`
	Status  string                       `json:"status"`
	Results datastructs.GetAccountResult `json:"results"`
}

func performPortfolioRequest(req *http.Request) (*getPortfolioResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj getPortfolioResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}

	return &respObj, nil
}

//  helper functions

func performAccountRequest(req *http.Request) (*getAccountResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj getAccountResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}

	return &respObj, nil
}
