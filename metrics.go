package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("pushing mtreics to prometheus")
	//prometheus data use grafan or prometheus here
	return s.next.FetchPrice(ctx, ticker)
}
