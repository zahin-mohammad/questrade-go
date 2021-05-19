package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TODO: Doc for parameters

func NewQuestradeAPIClient(
	ctx context.Context,
	isTest bool,
	refreshToken string,
	retryTimeInSeconds *time.Duration,
	maxRetry *int,
	checkRetry ...func(context.Context, *http.Response, error) (bool, error),
) (*QuestradeAPIClient, error) {
	// Inject custom http client via context
	ctx = context.WithValue(
		ctx, oauth2.HTTPClient,
		NewRetryRateLimitClient(retryTimeInSeconds, maxRetry, checkRetry...))
	questradeAPIClient := QuestradeAPIClient{ctx: ctx, isTest: isTest, refreshToken: refreshToken}
	tokenSource := oauth2.ReuseTokenSource(nil, &questradeAPIClient)
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	apiServer, ok := token.Extra(apiServerKey).(string)
	if !ok {
		return nil, fmt.Errorf("%s key was not set in token.Extra", apiServerKey)
	}
	questradeAPIClient.apiURL = apiServer
	questradeAPIClient.Client = oauth2.NewClient(ctx, tokenSource)
	return &questradeAPIClient, nil
}

func (client *QuestradeAPIClient) GetAccounts() (*AccountsResponse, error) {
	url := fmt.Sprintf("%s%s%s", client.apiURL, version, getAccounts)
	body, err := client.doRequest(url)
	if err != nil {
		return nil, err
	}
	var accountsResponse AccountsResponse
	if err = json.Unmarshal(body, &accountsResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}
	return &accountsResponse, nil
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
		log.Println(string(body))
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
	var accountPositionsResponse AccountPositionsResponse
	if err = json.Unmarshal(body, &accountPositionsResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}
	return &accountPositionsResponse, nil
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
	var accountActivitiesResponse AccountActivitiesResponse
	if err = json.Unmarshal(body, &accountActivitiesResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}
	return &accountActivitiesResponse, nil
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
	var accountExecutionsResponse AccountExecutionsResponse
	if err = json.Unmarshal(body, &accountExecutionsResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}
	return &accountExecutionsResponse, nil
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
		params.Add("ids", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(orderIDs)), ","), "[]"))
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
	var accountOrdersResponse AccountOrdersResponse
	if err = json.Unmarshal(body, &accountOrdersResponse); err != nil {
		log.Println(string(body))
		return nil, err
	}
	return &accountOrdersResponse, nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Inspired by: https://github.com/arianitu/go-questrade-oauth2/blob/master/questrade.go
//////////////////////////////////////////////////////////////////////////////////////////

func (client *QuestradeAPIClient) Token() (*oauth2.Token, error) {
	oauthClient := oauth2.NewClient(client.ctx, nil)
	apiURL := oauth2URL
	if client.isTest {
		apiURL = oauth2URLTest
	}
	resp, err := oauthClient.Get(apiURL + client.refreshToken)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New("Invalid Refresh Token")
	}
	var authResp authResponse
	// TODO: Clean closure
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &authResp); err != nil {
		return nil, err
	}
	token := &oauth2.Token{
		AccessToken: authResp.AccessToken,
		TokenType:   authResp.TokenType,
	}
	extra := url.Values{}
	extra.Add(apiServerKey, authResp.ApiServer)
	token = token.WithExtra(extra)
	if secs := authResp.ExpiresIn; secs > 0 {
		token.Expiry = time.Now().Add(time.Duration(secs) * time.Second)
	}
	client.refreshToken = authResp.RefreshToken
	fmt.Println(client.refreshToken)
	return token, nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Helpers
//////////////////////////////////////////////////////////////////////////////////////////

func (client *QuestradeAPIClient) doRequest(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
