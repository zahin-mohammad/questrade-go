package app

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (client *QuestradeAPIClient) GetMarkets() (*MarketsResponse, error) {
	url := fmt.Sprintf("%s%s%s", client.apiURL, version, getMarkets)
	body, err := client.doRequest(url)
	if err != nil {
		return nil, err
	}
	var resp MarketsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetMarketsQuotes(
	id *int,
	ids ...int,
) (*MarketQuotesResponse, error) {
	urlString := ""
	if id != nil {
		urlString = fmt.Sprintf("%s%s%s/%d", client.apiURL, version, getMarketQuotes, *id)
	} else {
		params := url.Values{}
		params.Add("ids", intSliceToString(",", ids...))
		urlString = fmt.Sprintf("%s%s%s?%s", client.apiURL, version, getMarketQuotes, params.Encode())
	}
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var resp MarketQuotesResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
