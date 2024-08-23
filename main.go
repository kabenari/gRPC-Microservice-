package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	price, err := svc.FetchPrice(context.Background(), "USDTC")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(price)
}
