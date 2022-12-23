package usecase

import (
	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/log"
	"github.com/rymiyamoto/affiliate-api/service"
)

type (
	IRakuten interface {
		Exec(f *flag.Affiliate) error
	}

	Rakuten struct {
		rakutenClientService service.IRakutenClient
	}
)

// NewRakuten initialize
func NewRakuten() IRakuten {
	return &Rakuten{
		rakutenClientService: service.NewRakutenClient(),
	}
}

// Exec 実行
func (u *Rakuten) Exec(f *flag.Affiliate) error {
	log.Info("rakuten")
	if err := u.rakutenClientService.GetProduct(); err != nil {
		return err
	}
	return nil
}
