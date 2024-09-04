package main

import (
	"context"
	"flag"
	"fmt"
	"grpc/client"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	log.Println("Starting the service...")
	c := client.New("http://localhost:3000")
	log.Println("Client initialized.")

	price, err := c.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal("Error fetching price:", err)
	}

	log.Println("Price fetched successfully.")
	fmt.Printf("Fetched Price: %+v\n", price)

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	service := newJSONAPIServer(*listenAddr, svc)
	log.Println("Service is about to run...")

	service.Run()

	log.Println("Service has started.")
}
