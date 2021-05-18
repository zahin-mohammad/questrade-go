package app

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type QuestradeAPIClient struct {
	ctx          context.Context
	isTest       bool
	refreshToken string
	ApiServerURL string
	*http.Client
}

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
	questradeAPIClient.ApiServerURL = apiServer
	questradeAPIClient.Client = oauth2.NewClient(ctx, tokenSource)
	return &questradeAPIClient, nil
}
