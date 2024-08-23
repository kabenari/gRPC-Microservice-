package main

import (
	"context"
	"fmt"
)

// pricefetcher is a interface that can fetch a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher strcut impolemetns the PriceFetcher interface
type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC":   20.000,
	"USDTC": 1090.90,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("th give nticker is unsupported")
	}

	return price, nil
}
