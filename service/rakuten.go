package service

import (
	"fmt"
	"time"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/repository"
)

type (
	IRakuten interface {
		Exec(f *flag.Affiliate) error
	}

	Rakuten struct {
		rakutenClientRepository repository.IRakutenClient
	}
)

// NewRakuten initialize
func NewRakuten() IRakuten {
	return &Rakuten{
		rakutenClientRepository: repository.NewRakutenClient(),
	}
}

// Exec 実行
func (s *Rakuten) Exec(f *flag.Affiliate) error {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if err := s.rakutenClientRepository.GetProduct(); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
