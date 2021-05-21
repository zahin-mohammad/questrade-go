package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/oauth2"
	"golang.org/x/time/rate"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// TODO: Doc for parameters

func NewQuestradeAPIClient(
	ctx context.Context,
	isTest bool,
	refreshToken string,
	retryTimeInSeconds *time.Duration, // optional
	maxRetry *int, // optional
	checkRetry ...func(context.Context, *http.Response, error) (bool, error), // optional
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

// NewRetryRateLimitClient : a rate limited retry client.
// All parameters are optional. If nil/zero values are passed in default values will be used.
func NewRetryRateLimitClient(
	retryTimeInSeconds *time.Duration,
	maxRetry *int,
	checkRetry ...func(context.Context, *http.Response, error) (bool, error),
) *http.Client {
	// https://www.questrade.com/api/documentation/rate-limiting
	rateLimiter := rate.NewLimiter(rate.Every(time.Second/20), 1)
	retryClient := retryablehttp.NewClient()
	if len(checkRetry) == 1 {
		retryClient.CheckRetry = checkRetry[0]
	} else {
		retryClient.CheckRetry = DefaultCheckRetry
	}
	if retryTimeInSeconds != nil {
		retryClient.RetryWaitMin = *retryTimeInSeconds
	} else {
		retryClient.RetryWaitMin = defaultRetryMin
	}
	if maxRetry != nil {
		retryClient.RetryMax = *maxRetry
	} else {
		retryClient.RetryMax = defaultMaxRetry
	}
	retryClient.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, retry int) {
		if err := rateLimiter.Wait(context.Background()); err != nil {
			log.Printf("ERROR WAITING FOR LIMIT: %s\n", err.Error())
			return
		}
	}
	return retryClient.StandardClient()
}

func DefaultCheckRetry(
	_ context.Context, resp *http.Response, err error,
) (bool, error) {
	// Don't retry non-get requests
	if resp.Request.Method != GET && resp.Request.Method != "" {
		return false, nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		if err != nil {
			log.Printf("retry error: StatusCode %d Error %s\n", resp.StatusCode, err.Error())
		} else {
			log.Printf("retry error: StatusCode %d\n", resp.StatusCode)
		}
		requestDump, err := httputil.DumpRequest(resp.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		// TODO: Remove these prints
		fmt.Println(string(requestDump))
		responseDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(responseDump))
		return true, err
	}
	return false, nil
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
	// TODO: Remove dump
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(responseDump))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
