package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	ENV_FILE                    = "../.env"
	QUESTRADE_REFRESH_TOKEN_KEY = "QUESTRADE_REFRESH_TOKEN"
	ACCOUNT_ID_NUMBER           = "ACCOUNT_ID_NUMBER"
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

	// TEARDOWN
	env[QUESTRADE_REFRESH_TOKEN_KEY] = questradeAPIClient.refreshToken
	_ = godotenv.Write(env, ENV_FILE)
}
