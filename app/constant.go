package app

import "time"

const (
	version       = "v1/"
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
	getAccountOrders     = "accounts/%s/orders" //accountID, comma delimited orderID's
	getAccounts          = "accounts"
	getAccountBalances   = "accounts/%s/balances"   //accountID
	getAccountPositions  = "accounts/%s/positions"  //accountID
	getAccountActivities = "accounts/%s/activities" //accountID
	getAccountExecutions = "accounts/%s/executions" //accountID

)

// Market Endpoints
const (
	getMarkets               = "markets"
	getMarketQuotes          = "markets/quotes"     //symbolID
	GetMarketCandles         = "markets/candles/%s" //symbolID
	GetMarketQuoteStrategies = "markets/quotes/strategies"
	GetMarketQuoteOptions    = "markets/quotes/options"
)

const (
	getSymbolsSearch = "symbols/search"
	getSymbols       = "symbols"            //symboldID
	getSymbolOptions = "symbols/%d/options" //symbolID

)
