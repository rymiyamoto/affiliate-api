package usecase

import (
	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/log"
)

type (
	IYahoo interface {
		Exec(f *flag.Affiliate) error
	}

	Yahoo struct {
	}
)

// NewYahoo initialize
func NewYahoo() IYahoo {
	return &Yahoo{}
}

// Exec 実行
func (u *Yahoo) Exec(f *flag.Affiliate) error {
	log.Info("yahoo")
	return nil
}
