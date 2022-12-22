package flag

import (
	"github.com/rymiyamoto/affiliate-api/model/common"
	"github.com/urfave/cli/v2"
)

type (
	Affiliate struct {
		ShopType common.ShopType
		Code     string
	}
)

func NewAffiliate(ctx *cli.Context) *Affiliate {
	return &Affiliate{
		ShopType: common.ToShopType(ctx.String("shop")),
		Code:     ctx.String("code"),
	}
}
