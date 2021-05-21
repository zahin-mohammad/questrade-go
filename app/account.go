package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

func (client *QuestradeAPIClient) GetAccounts() (*AccountsResponse, error) {
	url := fmt.Sprintf("%s%s%s", client.apiURL, version, getAccounts)
	body, err := client.doRequest(url)
	if err != nil {
		return nil, err
	}
	var resp AccountsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetAccountBalances(
	accountID string,
) (*AccountBalancesResponse, error) {
	uri := fmt.Sprintf(getAccountBalances, accountID)
	urlString := fmt.Sprintf("%s%s%s", client.apiURL, version, uri)
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var accountBalancesResponse AccountBalancesResponse
	if err = json.Unmarshal(body, &accountBalancesResponse); err != nil {
		return nil, err
	}
	return &accountBalancesResponse, nil
}

func (client *QuestradeAPIClient) GetAccountPositions(
	accountID string,
) (*AccountPositionsResponse, error) {
	uri := fmt.Sprintf(getAccountPositions, accountID)
	urlString := fmt.Sprintf("%s%s%s", client.apiURL, version, uri)
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var resp AccountPositionsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetAccountActivities(
	accountID string,
	startTime time.Time,
	endTime time.Time,
) (*AccountActivitiesResponse, error) {
	uri := fmt.Sprintf(getAccountActivities, accountID)
	params := url.Values{}
	params.Add("startTime", startTime.Format(time.RFC3339))
	params.Add("endTime", endTime.Format(time.RFC3339))
	urlString := fmt.Sprintf("%s%s%s?%s", client.apiURL, version, uri, params.Encode())
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var resp AccountActivitiesResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetAccountExecutions(
	accountID string,
	startTime time.Time,
	endTime time.Time,
) (*AccountExecutionsResponse, error) {
	uri := fmt.Sprintf(getAccountExecutions, accountID)
	params := url.Values{}
	params.Add("startTime", startTime.Format(time.RFC3339))
	params.Add("endTime", endTime.Format(time.RFC3339))
	urlString := fmt.Sprintf("%s%s%s?%s", client.apiURL, version, uri, params.Encode())
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var resp AccountExecutionsResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *QuestradeAPIClient) GetAccountOrders(
	accountID string,
	startTime time.Time,
	endTime time.Time,
	stateFilter OrderStateENUM,
	orderIDs ...int,
) (*AccountOrdersResponse, error) {
	uri := fmt.Sprintf(getAccountOrders, accountID)
	params := url.Values{}
	if len(orderIDs) > 0 {
		params.Add("ids", intSliceToString(",", orderIDs...))
	} else {
		params.Add("startTime", startTime.Format(time.RFC3339))
		params.Add("endTime", endTime.Format(time.RFC3339))
		params.Add("stateFilter", string(stateFilter))
	}
	urlString := fmt.Sprintf("%s%s%s?%s", client.apiURL, version, uri, params.Encode())
	body, err := client.doRequest(urlString)
	if err != nil {
		return nil, err
	}
	var resp AccountOrdersResponse
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
