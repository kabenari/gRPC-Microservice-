package client

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc/types"
	"net/http"
)

type client struct {
	endpoint string
}

func New(endpoint string) *client {
	return &client{
		endpoint: endpoint,
	}
}

func (c *client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service respond with non ok status code:%s", httpErr["err"])
	}

	return priceResp, nil
}
