package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	ENV_FILE                    = "../.env"
	QUESTRADE_REFRESH_TOKEN_KEY = "QUESTRADE_REFRESH_TOKEN"
	ACCOUNT_ID_NUMBER           = "ACCOUNT_ID_NUMBER"
	BMO_SYMBOL_ID               = 9292
	AAPL_SYMBOL_ID              = 8049
	BMOOF_SYMBOL_ID             = 35518868
)

// Testing done with prod because their test app is unstable...

func TestNewOauthClient(t *testing.T) {
	// SETUP
	env, _ := godotenv.Read(ENV_FILE)
	refreshToken := env[QUESTRADE_REFRESH_TOKEN_KEY]
	accountID := env[ACCOUNT_ID_NUMBER]
	questradeAPIClient, err := NewQuestradeAPIClient(
		context.Background(), false,
		refreshToken,
		nil, nil)
	assert.NoError(t, err)
	env[QUESTRADE_REFRESH_TOKEN_KEY] = questradeAPIClient.refreshToken
	_ = godotenv.Write(env, ENV_FILE)

	t.Run("Initialization", func(t *testing.T) {
		assert.NotEmpty(t, questradeAPIClient.refreshToken)
	})

	//////////////////////////////////////////////////////////////////////////////////////////
	// Accounts
	//////////////////////////////////////////////////////////////////////////////////////////

	t.Run("GetAccounts", func(t *testing.T) {
		resp, err := questradeAPIClient.GetAccounts()
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Accounts)
		//fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetAccountBalances", func(t *testing.T) {
		resp, err := questradeAPIClient.GetAccountBalances(accountID)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.CombinedBalances)
		//fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetAccountPositions", func(t *testing.T) {
		resp, err := questradeAPIClient.GetAccountPositions(accountID)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Positions)
		//fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetAccountActivities", func(t *testing.T) {
		endTime := time.Now()
		startTime := endTime.Add(-time.Hour * 24 * 30)
		resp, err := questradeAPIClient.GetAccountActivities(accountID, startTime, endTime)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Activities)
		//fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetAccountExecutions", func(t *testing.T) {
		endTime := time.Now()
		startTime := endTime.Add(-time.Hour * 24 * 30)
		resp, err := questradeAPIClient.GetAccountExecutions(accountID, startTime, endTime)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Executions)
		//fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetAccountOrders", func(t *testing.T) {
		endTime := time.Now()
		startTime := endTime.Add(-time.Hour * 24 * 30)
		resp, err := questradeAPIClient.GetAccountOrders(accountID, startTime, endTime, OrderStateAll)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Orders)
		orders := []int{}
		for _, order := range resp.Orders {
			orders = append(orders, order.ID)
		}
		respFromIDS, err := questradeAPIClient.GetAccountOrders(accountID, startTime, endTime, OrderStateAll, orders...)
		assert.NoError(t, err)
		assert.Equal(t, len(resp.Orders), len(respFromIDS.Orders))
		//fmt.Println(PrettyJSON(resp))
	})

	//////////////////////////////////////////////////////////////////////////////////////////
	// Markets
	//////////////////////////////////////////////////////////////////////////////////////////

	t.Run("GetMarkets", func(t *testing.T) {
		resp, err := questradeAPIClient.GetMarkets()
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Markets)
		fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetMarketsQuotes", func(t *testing.T) {
		symb := BMOOF_SYMBOL_ID
		resp, err := questradeAPIClient.GetMarketsQuotes(&symb)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Quotes)
		fmt.Println(PrettyJSON(resp))

		resp, err = questradeAPIClient.GetMarketsQuotes(nil, BMO_SYMBOL_ID, BMOOF_SYMBOL_ID)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Quotes)
		fmt.Println(PrettyJSON(resp))
	})

	//////////////////////////////////////////////////////////////////////////////////////////
	// Symbol
	//////////////////////////////////////////////////////////////////////////////////////////

	t.Run("GetSymbolOptions", func(t *testing.T) {
		resp, err := questradeAPIClient.GetSymbolOptions(BMO_SYMBOL_ID)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.OptionChain)
		fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetSymbolsSearch", func(t *testing.T) {
		resp, err := questradeAPIClient.GetSymbolsSearch("BMO")
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Symbols)
		fmt.Println(PrettyJSON(resp))
	})

	t.Run("GetSymbols", func(t *testing.T) {
		symb := AAPL_SYMBOL_ID
		resp, err := questradeAPIClient.GetSymbols(&symb)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Symbols)
		fmt.Println(PrettyJSON(resp))

		resp, err = questradeAPIClient.GetSymbols(nil, AAPL_SYMBOL_ID, BMO_SYMBOL_ID)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Symbols)
		fmt.Println(PrettyJSON(resp))
	})

	// TEARDOWN
	env[QUESTRADE_REFRESH_TOKEN_KEY] = questradeAPIClient.refreshToken
	_ = godotenv.Write(env, ENV_FILE)
}
