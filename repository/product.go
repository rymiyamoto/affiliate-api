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

	switch shop {
	case common.ShopTypeRakuten:
		products := model.Products{
			{
				ID:       10,
				ShopType: common.ShopTypeRakuten,
				Code:     "",
				Price:    8000,
			},
			{
				ID:       11,
				ShopType: common.ShopTypeRakuten,
				Code:     "",
				Price:    5000,
			},
			{
				ID:       12,
				ShopType: common.ShopTypeRakuten,
				Code:     "",
				Price:    300,
			},
		}
		ms = append(ms, products...)
	case common.ShopTypeYahoo:
		products := model.Products{
			{
				ID:       20,
				ShopType: common.ShopTypeYahoo,
				Code:     "",
				Price:    7000,
			},
			{
				ID:       21,
				ShopType: common.ShopTypeYahoo,
				Code:     "",
				Price:    4000,
			},
			{
				ID:       22,
				ShopType: common.ShopTypeYahoo,
				Code:     "",
				Price:    500,
			},
		}
		ms = append(ms, products...)
	default:
		return nil, fmt.Errorf("not define type. shop_type: %d", shop)
	}

	return &ms, nil
}
