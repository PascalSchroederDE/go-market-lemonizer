package trading

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	api_helper "github.com/PascalSchroederDE/go-market-lemonizer/internal/lemonizer/api"
	datastructs "github.com/PascalSchroederDE/go-market-lemonizer/pkg/lemonizer/datastructs"
)

// Define API JSON response to Go Object

type placeOrderResponse struct {
	Time    time.Time                     `json:"time"`
	Mode    string                        `json:"mode"`
	Status  string                        `json:"status"`
	Results datastructs.PlaceOrderResults `json:"results"`
}

type activateOrderResponse struct {
	Time   time.Time `json:"time"`
	Mode   string    `json:"mode"`
	Status string    `json:"status"`
}

type deleteOrderResponse struct {
	Time   time.Time `json:"time"`
	Mode   string    `json:"mode"`
	Status string    `json:"status"`
}

//  helper functions

func performPlaceOrderRequest(req *http.Request) (*placeOrderResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj placeOrderResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}

	return &respObj, nil
}

func performActivateOrderRequest(req *http.Request) (*activateOrderResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj activateOrderResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}

	return &respObj, nil
}

func performDeleteOrderRequest(req *http.Request) (*deleteOrderResponse, error) {
	respBody, err := api_helper.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	var respObj deleteOrderResponse
	if err := json.Unmarshal(respBody, &respObj); err != nil {
		return nil, errors.New("unexpected response form")
	}

	return &respObj, nil
}
