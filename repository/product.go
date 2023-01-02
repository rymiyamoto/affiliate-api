package repository

import (
	"fmt"

	"github.com/rymiyamoto/affiliate-api/model"
	"github.com/rymiyamoto/affiliate-api/model/common"
)

type (
	// IProduct interface
	IProduct interface {
		ByShopType(shop common.ShopType) (*model.Products, error)
	}

	// Product struct
	Product struct {
	}
)

func NewProduct() IProduct {
	return &Product{}
}

// ByShopType ショップ種別から情報取得
func (r *Product) ByShopType(shop common.ShopType) (*model.Products, error) {
	ms := model.Products{}

	//nolint:exhaustive
	switch shop {
	case common.ShopTypeRakuten:
		products := model.Products{
			{
				ID:       10, //nolint:gomnd
				ShopType: common.ShopTypeRakuten,
				Code:     "8e27892c206a7118a74445c5a0825a41",
				Price:    8000, //nolint:gomnd
			},
			{
				ID:       11, //nolint:gomnd
				ShopType: common.ShopTypeRakuten,
				Code:     "a3da4c917cdfdbce5577f6923533f003",
				Price:    5000, //nolint:gomnd
			},
			{
				ID:       12, //nolint:gomnd
				ShopType: common.ShopTypeRakuten,
				Code:     "1c3d37510cfe620aa200862886a6502c",
				Price:    300, //nolint:gomnd
			},
		}
		ms = append(ms, products...)
	case common.ShopTypeYahoo:
		products := model.Products{
			{
				ID:       20, //nolint:gomnd
				ShopType: common.ShopTypeYahoo,
				Code:     "",
				Price:    7000, //nolint:gomnd
			},
			{
				ID:       21, //nolint:gomnd
				ShopType: common.ShopTypeYahoo,
				Code:     "",
				Price:    4000, //nolint:gomnd
			},
			{
				ID:       22, //nolint:gomnd
				ShopType: common.ShopTypeYahoo,
				Code:     "",
				Price:    500, //nolint:gomnd
			},
		}
		ms = append(ms, products...)
	default:
		return nil, fmt.Errorf("not define type. shop_type: %d", shop)
	}

	return &ms, nil
}
