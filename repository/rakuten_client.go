package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rymiyamoto/affiliate-api/util/env"
	"github.com/rymiyamoto/affiliate-api/util/log"
)

type (
	IRakutenClient interface {
		GetProduct(code string) error
	}

	RakutenClient struct {
	}
)

const (
	rakutenProductSearchEndpoint = "https://app.rakuten.co.jp/services/api/Product/Search/20170426?format=json"
)

// NewRakutenClient initialize
func NewRakutenClient() IRakutenClient {
	return &RakutenClient{}
}

// GetProduct 商品取得
func (s *RakutenClient) GetProduct(code string) error {
	//nolint
	req, err := http.NewRequest("GET", rakutenProductSearchEndpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to make rakuten product search request. err: %w", err)
	}

	appID, err := env.Value("RAKUTEN_APP_ID")
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("applicationId", appID)
	q.Add("productId", code)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("failed to request rakuten product search. err: %w", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to perse response rakuten product search. err: %w", err)
	}

	log.Infof("body: %s", string(body))

	return nil
}
