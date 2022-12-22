package dispatch

import (
	"fmt"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/log"
	"github.com/rymiyamoto/affiliate-api/model/common"
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

	f := flag.NewAffiliate(ctx)

	switch f.ShopType {
	case common.ShopTypeRakuten:
		fmt.Println("rakuten")
	case common.ShopTypeYahoo:
		fmt.Println("yahoo")
	default:
		return fmt.Errorf("not found shop. f: %+v", f)
	}

	return nil
}
