package app

import (
	"context"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

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
		return true, err
	}
	return false, nil
}
