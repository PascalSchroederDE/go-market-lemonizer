package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func BuildGetRequestWithParameter(apiKey string, baseUrl string, path string, parameter map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseUrl+path, nil)
	if err != nil {
		return nil, errors.New("create request error")
	}

	q := req.URL.Query()
	for key, value := range parameter {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+apiKey)

	return req, nil
}

func BuildGetRequestWithoutParameter(apiKey string, baseUrl string, path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseUrl+path, nil)

	if err != nil {
		return nil, errors.New("create request error")
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)

	return req, nil
}

func BuildPostRequestWithParameter(apiKey string, baseUrl string, path string, bodyValues map[string]string) (*http.Request, error) {
	body := url.Values{}
	for key, value := range bodyValues {
		body.Set(key, value)
	}

	req, err := http.NewRequest("POST", baseUrl+path, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, errors.New("create request error")
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	return req, nil
}

func BuildPostRequestWithoutParameter(apiKey string, baseUrl string, path string) (*http.Request, error) {
	req, err := http.NewRequest("POST", baseUrl+path, nil)
	if err != nil {
		return nil, errors.New("create request error")
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)

	return req, nil
}

func BuildDeleteRequestWithoutParameter(apiKey string, baseUrl string, path string) (*http.Request, error) {
	req, err := http.NewRequest("DEL", baseUrl+path, nil)
	if err != nil {
		return nil, errors.New("create request error")
	}

	req.Header.Add("Authorization", "Bearer "+apiKey)

	return req, nil
}

func PerformRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("api request error")
	}

	defer response.Body.Close()
	respBody, _ := ioutil.ReadAll(response.Body)

	return respBody, nil
}
