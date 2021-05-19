package app

import "time"

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type authResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	ApiServer    string `json:"api_server"`
}

type AccountActivitiesResponse struct {
	Activities []struct {
		TradeDate       time.Time `json:"tradeDate"`
		TransactionDate time.Time `json:"transactionDate"`
		SettlementDate  time.Time `json:"settlementDate"`
		Action          string    `json:"action"`
		Symbol          string    `json:"symbol"`
		SymbolId        int       `json:"symbolId"`
		Description     string    `json:"description"`
		Currency        string    `json:"currency"`
		Quantity        int       `json:"quantity"`
		Price           float64   `json:"price"`
		GrossAmount     float64   `json:"grossAmount"`
		Commission      float64   `json:"commission"`
		NetAmount       float64   `json:"netAmount"`
		Type            string    `json:"type"`
	} `json:"activities"`
}

type OrderStateENUM string

const (
	OrderStateAll    = OrderStateENUM("All")
	OrderStateOpen   = OrderStateENUM("Open")
	OrderStateClosed = OrderStateENUM("Closed")
)

type AccountOrdersResponse struct {
	Orders []struct {
		ID                       int            `json:"id"`
		Symbol                   string         `json:"symbol"`
		SymbolID                 int            `json:"symbolId"`
		TotalQuantity            int            `json:"totalQuantity"`
		OpenQuantity             int            `json:"openQuantity"`
		FilledQuantity           int            `json:"filledQuantity"`
		CanceledQuantity         int            `json:"canceledQuantity"`
		Side                     string         `json:"side"`
		Type                     string         `json:"type"`
		LimitPrice               float64        `json:"limitPrice"`
		StopPrice                float64        `json:"stopPrice"`
		IsAllOrNone              bool           `json:"isAllOrNone"`
		IsAnonymous              bool           `json:"isAnonymous"`
		IcebergQty               int            `json:"icebergQty"`
		MinQuantity              int            `json:"minQuantity"`
		AvgExecPrice             float64        `json:"avgExecPrice"`
		LastExecPrice            float64        `json:"lastExecPrice"`
		Source                   string         `json:"source"`
		TimeInForce              string         `json:"timeInForce"`
		GtdDate                  time.Time      `json:"gtdDate"`
		State                    string         `json:"state"`
		ClientReasonStr          string         `json:"clientReasonStr"`
		ChainID                  int            `json:"chainId"`
		CreationTime             time.Time      `json:"creationTime"`
		UpdateTime               time.Time      `json:"updateTime"`
		Notes                    string         `json:"notes"`
		PrimaryRoute             string         `json:"primaryRoute"`
		SecondaryRoute           string         `json:"secondaryRoute"`
		OrderRoute               string         `json:"orderRoute"`
		VenueHoldingOrder        string         `json:"venueHoldingOrder"`
		ComissionCharged         float64        `json:"comissionCharged"`
		ExchangeOrderID          string         `json:"exchangeOrderId"`
		IsSignificantShareHolder bool           `json:"isSignificantShareHolder"`
		IsInsider                bool           `json:"isInsider"`
		IsLimitOffsetInDollar    bool           `json:"isLimitOffsetInDollar"`
		UserID                   int            `json:"userId"`
		PlacementCommission      float64        `json:"placementCommission"`
		Legs                     []interface{}  `json:"legs"`
		StrategyType             string         `json:"strategyType"`
		TriggerStopPrice         float64        `json:"triggerStopPrice"`
		OrderGroupID             int            `json:"orderGroupId"`
		OrderClass               OrderClassEnum `json:"orderClass"`
		MainChainID              int            `json:"mainChainId"`
	} `json:"orders"`
}

type OrderClassEnum string

const (
	OrderClassEmpty   = OrderClassEnum("")
	OrderClassPrimary = OrderClassEnum("Primary")
	OrderClassProfit  = OrderClassEnum("Profit")
	OrderClassLoss    = OrderClassEnum("Loss")
)

type AccountExecutionsResponse struct {
	Executions []struct {
		Symbol                   string    `json:"symbol"`
		SymbolID                 int       `json:"symbolId"`
		Quantity                 int       `json:"quantity"`
		Side                     string    `json:"side"`
		Price                    float64   `json:"price"`
		ID                       int       `json:"id"`
		OrderID                  int       `json:"orderId"`
		OrderChainID             int       `json:"orderChainId"`
		ExchangeExecID           string    `json:"exchangeExecId"`
		Timestamp                time.Time `json:"timestamp"`
		Notes                    string    `json:"notes"`
		Venue                    string    `json:"venue"`
		TotalCost                float64   `json:"totalCost"`
		OrderPlacementCommission float64   `json:"orderPlacementCommission"`
		Commission               float64   `json:"commission"`
		ExecutionFee             float64   `json:"executionFee"`
		SecFee                   float64   `json:"secFee"`
		CanadianExecutionFee     float64   `json:"canadianExecutionFee"`
		ParentID                 int       `json:"parentId"`
	} `json:"executions"`
}

type AccountBalancesResponse struct {
	PerCurrencyBalances    []Balance `json:"perCurrencyBalances"`
	CombinedBalances       []Balance `json:"combinedBalances"`
	SodPerCurrencyBalances []Balance `json:"sodPerCurrencyBalances"`
	SodCombinedBalances    []Balance `json:"sodCombinedBalances"`
}

type Balance struct {
	Currency          string  `json:"currency"`
	Cash              float64 `json:"cash"`
	MarketValue       float64 `json:"marketValue"`
	TotalEquity       float64 `json:"totalEquity"`
	BuyingPower       float64 `json:"buyingPower"`
	MaintenanceExcess float64 `json:"maintenanceExcess"`
	IsRealTime        bool    `json:"isRealTime"`
}

type AccountPositionsResponse struct {
	Positions []struct {
		Symbol             string  `json:"symbol"`
		SymbolId           int     `json:"symbolId"`
		OpenQuantity       int     `json:"openQuantity"`
		CurrentMarketValue float64 `json:"currentMarketValue"`
		CurrentPrice       float64 `json:"currentPrice"`
		AverageEntryPrice  float64 `json:"averageEntryPrice"`
		ClosedPnl          float64 `json:"closedPnl"`
		OpenPnl            float64 `json:"openPnl"`
		TotalCost          float64 `json:"totalCost"`
		IsRealTime         bool    `json:"isRealTime"`
		IsUnderReorg       bool    `json:"isUnderReorg"`
	} `json:"positions"`
}

type AccountsResponse struct {
	Accounts []struct {
		Type              string `json:"type"`
		Number            string `json:"number"`
		Status            string `json:"status"`
		IsPrimary         bool   `json:"isPrimary"`
		IsBilling         bool   `json:"isBilling"`
		ClientAccountType string `json:"clientAccountType"`
	} `json:"accounts"`
}

/////////////////////////////////////////////////////////////////////////////////////////////
// Market
/////////////////////////////////////////////////////////////////////////////////////////////

type MarketCandlesResponse struct {
	Candles []struct {
		Start  time.Time `json:"start"`
		End    time.Time `json:"end"`
		Low    float64   `json:"low"`
		High   float64   `json:"high"`
		Open   float64   `json:"open"`
		Close  float64   `json:"close"`
		Volume int       `json:"volume"`
	} `json:"candles"`
}

type MarketStrategyQuotesResponse struct {
	StategyQuotes []struct {
		VariantID    int     `json:"variantId"`
		BidPrice     float64 `json:"bidPrice"`
		AskPrice     float64 `json:"askPrice"`
		Underlying   string  `json:"underlying"`
		UnderlyingID int     `json:"underlyingId"`
		OpenPrice    float64 `json:"openPrice"`
		Volatility   int     `json:"volatility"`
		Delta        int     `json:"delta"`
		Gamma        int     `json:"gamma"`
		Theta        int     `json:"theta"`
		Vega         int     `json:"vega"`
		Rho          int     `json:"rho"`
		IsRealTime   bool    `json:"isRealTime"`
	} `json:"stategyQuotes"`
}
type MarketOptionQuotesResponse struct {
	OptionQuotes []struct {
		Underlying          string  `json:"underlying"`
		UnderlyingID        int     `json:"underlyingId"`
		Symbol              string  `json:"symbol"`
		SymbolID            int     `json:"symbolId"`
		BidPrice            float64 `json:"bidPrice"`
		BidSize             int     `json:"bidSize"`
		AskPrice            float64 `json:"askPrice"`
		AskSize             int     `json:"askSize"`
		LastTradePriceTrHrs float64 `json:"lastTradePriceTrHrs"`
		LastTradePrice      float64 `json:"lastTradePrice"`
		LastTradeSize       int     `json:"lastTradeSize"`
		LastTradeTick       string  `json:"lastTradeTick"`
		LastTradeTime       string  `json:"lastTradeTime"`
		Volume              int     `json:"volume"`
		OpenPrice           int     `json:"openPrice"`
		HighPricehighPrice  float64 `json:"highPricehighPrice"`
		LowPrice            int     `json:"lowPrice"`
		Volatility          float64 `json:"volatility"`
		Delta               float64 `json:"delta"`
		Gamma               float64 `json:"gamma"`
		Theta               float64 `json:"theta"`
		Vega                float64 `json:"vega"`
		Rho                 float64 `json:"rho"`
		OpenInterest        int     `json:"openInterest"`
		Delay               int     `json:"delay"`
		IsHalted            bool    `json:"isHalted"`
		VWAP                int     `json:"VWAP"`
	} `json:"optionQuotes"`
}

type MarketQuotesResponse struct {
	Quotes []struct {
		Symbol              string  `json:"symbol"`
		SymbolID            int     `json:"symbolId"`
		Tier                string  `json:"tier"`
		BidPrice            float64 `json:"bidPrice"`
		BidSize             int     `json:"bidSize"`
		AskPrice            float64 `json:"askPrice"`
		AskSize             int     `json:"askSize"`
		LastTradePriceTrHrs float64 `json:"lastTradePriceTrHrs"`
		LastTradePrice      float64 `json:"lastTradePrice"`
		LastTradeSize       int     `json:"lastTradeSize"`
		LastTradeTick       string  `json:"lastTradeTick"`
		LastTradeTime       string  `json:"lastTradeTime"`
		Volume              int     `json:"volume"`
		OpenPrice           float64 `json:"openPrice"`
		HighPrice           float64 `json:"highPrice"`
		LowPrice            float64 `json:"lowPrice"`
		Delay               int     `json:"delay"`
		IsHalted            bool    `json:"isHalted"`
	} `json:"quotes"`
}

type MarketsResponse struct {
	Markets []struct {
		Name                 string   `json:"name"`
		TradingVenues        []string `json:"tradingVenues"`
		DefaultTradingVenue  string   `json:"defaultTradingVenue"`
		PrimaryOrderRoutes   []string `json:"primaryOrderRoutes"`
		SecondaryOrderRoutes []string `json:"secondaryOrderRoutes"`
		Level1Feeds          []string `json:"level1Feeds"`
		ExtendedStartTime    string   `json:"extendedStartTime"`
		StartTime            string   `json:"startTime"`
		EndTime              string   `json:"endTime"`
		Currency             string   `json:"currency"`
		SnapQuotesLimit      int      `json:"snapQuotesLimit"`
	} `json:"markets"`
}

type MarketOptionsResponse struct {
	Options []struct {
		ExpiryDate         string `json:"expiryDate"`
		Description        string `json:"description"`
		ListingExchange    string `json:"listingExchange"`
		OptionExerciseType string `json:"optionExerciseType"`
		ChainPerRoot       []struct {
			Root                string `json:"root"`
			ChainPerStrikePrice []struct {
				StrikePrice  int `json:"strikePrice"`
				CallSymbolID int `json:"callSymbolId"`
				PutSymbolID  int `json:"putSymbolId"`
			} `json:"chainPerStrikePrice"`
			Multiplier int `json:"multiplier"`
		} `json:"chainPerRoot"`
	} `json:"options"`
}

type MarketSymbolSearchResponse struct {
	Symbol []struct {
		Symbol          string `json:"symbol"`
		SymbolID        int    `json:"symbolId"`
		Description     string `json:"description"`
		SecurityType    string `json:"securityType"`
		ListingExchange string `json:"listingExchange"`
		IsTradable      bool   `json:"isTradable"`
		IsQuotable      bool   `json:"isQuotable"`
		Currency        string `json:"currency"`
	} `json:"symbol"`
}

type MarketSymbolsResponse struct {
	Symbols []struct {
		Symbol                     string      `json:"symbol"`
		SymbolID                   int         `json:"symbolId"`
		PrevDayClosePrice          float64     `json:"prevDayClosePrice"`
		HighPrice52                float64     `json:"highPrice52"`
		LowPrice52                 float64     `json:"lowPrice52"`
		AverageVol3Months          int         `json:"averageVol3Months"`
		AverageVol20Days           int         `json:"averageVol20Days"`
		OutstandingShares          int64       `json:"outstandingShares"`
		Eps                        float64     `json:"eps"`
		Pe                         float64     `json:"pe"`
		Dividend                   float64     `json:"dividend"`
		Yield                      float64     `json:"yield"`
		ExDate                     string      `json:"exDate"`
		MarketCap                  int64       `json:"marketCap"`
		TradeUnit                  int         `json:"tradeUnit"`
		OptionType                 interface{} `json:"optionType"`
		OptionDurationType         interface{} `json:"optionDurationType"`
		OptionRoot                 string      `json:"optionRoot"`
		OptionContractDeliverables struct {
			Underlyings []interface{} `json:"underlyings"`
			CashInLieu  int           `json:"cashInLieu"`
		} `json:"optionContractDeliverables"`
		OptionExerciseType interface{} `json:"optionExerciseType"`
		ListingExchange    string      `json:"listingExchange"`
		Description        string      `json:"description"`
		SecurityType       string      `json:"securityType"`
		OptionExpiryDate   interface{} `json:"optionExpiryDate"`
		DividendDate       string      `json:"dividendDate"`
		OptionStrikePrice  interface{} `json:"optionStrikePrice"`
		IsTradable         bool        `json:"isTradable"`
		IsQuotable         bool        `json:"isQuotable"`
		HasOptions         bool        `json:"hasOptions"`
		MinTicks           []struct {
			Pivot   int     `json:"pivot"`
			MinTick float64 `json:"minTick"`
		} `json:"minTicks"`
		IndustrySector   string `json:"industrySector"`
		IndustryGroup    string `json:"industryGroup"`
		IndustrySubGroup string `json:"industrySubGroup"`
	} `json:"symbols"`
}
