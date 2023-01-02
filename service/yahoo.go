package service

import (
	"time"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/repository"
)

type (
	IYahoo interface {
		Exec(f *flag.Affiliate) error
	}

	Yahoo struct {
		productRepository     repository.IProduct
		yahooClientRepository repository.IYahooClient
	}
)

// NewYahoo initialize
func NewYahoo() IYahoo {
	return &Yahoo{
		productRepository:     repository.NewProduct(),
		yahooClientRepository: repository.NewYahooClient(),
	}
}

// Exec 実行
func (s *Yahoo) Exec(f *flag.Affiliate) error {
	products, err := s.productRepository.ByShopType(f.ShopType)
	if err != nil {
		return err
	}

	for _, v := range *products {
		if err := s.yahooClientRepository.GetProduct(v.Code); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
