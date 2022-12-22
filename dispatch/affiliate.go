package dispatch

import (
	"github.com/rymiyamoto/affiliate-api/log"
	"github.com/urfave/cli/v2"
)

type (
	IAffiliate interface {
		Run(ctx *cli.Context) error
	}

	Affiliate struct {
	}
)

// NewAffiliate initialize
func NewAffiliate() IAffiliate {
	return &Affiliate{}
}

// Run 本体
func (s *Affiliate) Run(ctx *cli.Context) error {
	log.Info("hello world")
	return nil
}
