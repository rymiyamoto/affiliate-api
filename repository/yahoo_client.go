package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rymiyamoto/affiliate-api/util/env"
	"github.com/rymiyamoto/affiliate-api/util/log"
)

type (
	IYahooClient interface {
		GetProduct(code string) error
	}

	YahooClient struct {
	}
)

const (
	yahooProductSearchEndpoint = "https://shopping.yahooapis.jp/ShoppingWebService/V3/itemSearch"
)

// NewYahooClient initialize
func NewYahooClient() IYahooClient {
	return &YahooClient{}
}

// GetProduct 商品取得
func (s *YahooClient) GetProduct(code string) error {
	//nolint
	req, err := http.NewRequest("GET", yahooProductSearchEndpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to make yahoo product search request. err: %w", err)
	}

	appID, err := env.Value("YAHOO_APP_ID")
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("appid", appID)
	q.Add("jan_code", code)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("failed to request yahoo product search. err: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to perse response yahoo product search. err: %w", err)
	}

	log.Infof("body: %s", string(body))

	return nil
}
