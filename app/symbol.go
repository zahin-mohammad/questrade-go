package app

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (client *QuestradeAPIClient) GetSymbolsSearch(
	prefix string,
) (*SymbolSearchResponse, error) {
	url := fmt.Sprintf("%s%s%s?prefix=%s", client.apiURL, version, getSymbolsSearch, prefix)
	body, err := client.doRequest(url)
	if err != nil {
		return nil, err
	}
	var resp SymbolSearchResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetSymbols(
	symbolID *int,
	symbolIDs ...int,
) (*SymbolsResponse, error) {
	urlString := ""
	if symbolID != nil {
		urlString = fmt.Sprintf("%s%s%s/%d", client.apiURL, version, getSymbols, *symbolID)
	} else {
		params := url.Values{}
		params.Add("ids", intSliceToString(",", symbolIDs...))
		urlString = fmt.Sprintf("%s%s%s?%s", client.apiURL, version, getSymbols, params.Encode())
	}
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var resp SymbolsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetSymbolOptions(
	symbolID int,
) (*SymbolOptionsResponse, error) {
	uri := fmt.Sprintf(getSymbolOptions, symbolID)
	url := fmt.Sprintf("%s%s%s", client.apiURL, version, uri)
	body, err := client.doRequest(url)
	if err != nil {
		return nil, err
	}
	var resp SymbolOptionsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
