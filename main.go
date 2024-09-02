package main

import "flag"

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen addres the service is runnnign")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	service := newJSONAPIServer(*listenAddr, svc)
	service.Run()
}
