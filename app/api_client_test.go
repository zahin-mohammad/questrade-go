package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ENV_FILE                    = "../.env"
	QUESTRADE_REFRESH_TOKEN_KEY = "QUESTRADE_REFRESH_TOKEN"
)

func TestNewOauthClient(t *testing.T) {
	env, err := godotenv.Read(ENV_FILE)
	assert.NoError(t, err)
	refreshToken := env[QUESTRADE_REFRESH_TOKEN_KEY]
	assert.NotEmpty(t, refreshToken)
	questradeAPIClient, err := NewQuestradeAPIClient(
		context.Background(), false,
		refreshToken,
		nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, questradeAPIClient)
	env[QUESTRADE_REFRESH_TOKEN_KEY] = questradeAPIClient.refreshToken
	_ = godotenv.Write(env, ENV_FILE)
}
