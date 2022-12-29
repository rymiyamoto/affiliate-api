package dispatch

import (
	"fmt"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/model/common"
	"github.com/rymiyamoto/affiliate-api/service"
	"github.com/urfave/cli/v2"
)

type (
	IAffiliate interface {
		Run(ctx *cli.Context) error
	}

	Affiliate struct {
		rakutenService service.IRakuten
		yahooService   service.IYahoo
	}
)

// NewAffiliate initialize
func NewAffiliate() IAffiliate {
	return &Affiliate{
		rakutenService: service.NewRakuten(),
		yahooService:   service.NewYahoo(),
	}
}

// Run 本体
func (d *Affiliate) Run(ctx *cli.Context) error {
	f := flag.NewAffiliate(ctx)

	switch f.ShopType {
	case common.ShopTypeRakuten:
		return d.rakutenService.Exec(f)
	case common.ShopTypeYahoo:
		return d.yahooService.Exec(f)
	default:
		return fmt.Errorf("not found shop. f: %+v", f)
	}
}
