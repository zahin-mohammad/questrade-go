package app

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)


type QuestradeAPIClient struct{
	ctx context.Context
	refreshToken string
	ApiServerURL string
	*http.Client
}

func NewOauthClient(
	ctx context.Context, refreshToken string,
) (*QuestradeAPIClient, error) {
	// Inject custom http client via context
	// TODO: Create retry client
	ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{Timeout: 2 * time.Second})
	questradeAPIClient := QuestradeAPIClient{ctx: ctx, refreshToken: refreshToken}
	tokenSource := oauth2.ReuseTokenSource(nil, questradeAPIClient)
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
