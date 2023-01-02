package service

import (
	"time"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/repository"
)

type (
	IRakuten interface {
		Exec(f *flag.Affiliate) error
	}

	Rakuten struct {
		productRepository       repository.IProduct
		rakutenClientRepository repository.IRakutenClient
	}
)

// NewRakuten initialize
func NewRakuten() IRakuten {
	return &Rakuten{
		productRepository:       repository.NewProduct(),
		rakutenClientRepository: repository.NewRakutenClient(),
	}
}

// Exec 実行
func (s *Rakuten) Exec(f *flag.Affiliate) error {
	products, err := s.productRepository.ByShopType(f.ShopType)
	if err != nil {
		return err
	}

	for _, v := range *products {
		if err := s.rakutenClientRepository.GetProduct(v.Code); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
