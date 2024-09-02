package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type PriceResponse struct {
	Ticker string  `json:ticker`
	Price  float64 `json:price`
}

type JSONServer struct {
	listenAddr string
	svc        PriceFetcher
}

func newJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONServer {
	return &JSONServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apifn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(100000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apifn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResp := PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJSON(w, http.StatusOK, &priceResp)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
