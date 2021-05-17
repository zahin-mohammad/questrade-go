package app

import "time"

const (
	oauth2URL     = "https://login.questrade.com/oauth2/token?grant_type=refresh_token&refresh_token="
	oauth2URLTest = "https://practicelogin.questrade.com/oauth2/token?grant_type=refresh_token&refresh_token="
	apiServerKey  = "api_server"
)

const (
	defaultRetryMin = time.Second * 10
	defaultMaxRetry = 3
)

const (
	GET = "GET"
)

// Account Endpoints
const (
	GetAccountActivities = "accounts/%s/activities"    //accountID
	GetAccountOrders     = "accounts/%s/orders?ids=%s" //accountID, comma delimited orderID's
	GetAccountExecutions = "accounts/%s/executions"    //accountID
	GetAccountBalances   = "accounts/%s/balances"      //accountID
	GetAccountPositions  = "accounts/%s/positions"     //accountID
	GetAccounts          = "accounts"
)

// Market Endpoints
const (
	GetMarketCandles         = "markets/candles/%s" //symbolID
	GetMarketQuoteStrategies = "markets/quotes/strategies"
	GetMarketQuoteOptions    = "markets/quotes/options"
	GetMarketQuotes          = "markets/quotes/%s" //symbolID
	GetMarkets               = "markets"
	GetMarketOptions         = "symbols/%s/options" //symbolID
	GetMarketSymbolSearch    = "symbols/search"
	GetMarketSymbols         = "symbols/%s" //symboldID
)

const ()
