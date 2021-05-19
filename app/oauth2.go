package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Inspired by: https://github.com/arianitu/go-questrade-oauth2/blob/master/questrade.go

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
