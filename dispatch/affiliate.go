package dispatch

import (
	"fmt"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/model/common"
	"github.com/rymiyamoto/affiliate-api/usecase"
	"github.com/urfave/cli/v2"
)

type (
	IAffiliate interface {
		Run(ctx *cli.Context) error
	}

	Affiliate struct {
		rakutenUsecase usecase.IRakuten
		yahooUsecase   usecase.IYahoo
	}
)

// NewAffiliate initialize
func NewAffiliate() IAffiliate {
	return &Affiliate{
		rakutenUsecase: usecase.NewRakuten(),
		yahooUsecase:   usecase.NewYahoo(),
	}
}

// Run 本体
func (d *Affiliate) Run(ctx *cli.Context) error {
	f := flag.NewAffiliate(ctx)

	switch f.ShopType {
	case common.ShopTypeRakuten:
		return d.rakutenUsecase.Exec(f)
	case common.ShopTypeYahoo:
		return d.yahooUsecase.Exec(f)
	default:
		return fmt.Errorf("not found shop. f: %+v", f)
	}
}
