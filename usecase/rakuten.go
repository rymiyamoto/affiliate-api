package usecase

import (
	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/log"
)

type (
	IRakuten interface {
		Exec(f *flag.Affiliate) error
	}

	Rakuten struct {
	}
)

// NewRakuten initialize
func NewRakuten() IRakuten {
	return &Rakuten{}
}

// Exec 実行
func (u *Rakuten) Exec(f *flag.Affiliate) error {
	log.Info("rakuten")
	return nil
}
