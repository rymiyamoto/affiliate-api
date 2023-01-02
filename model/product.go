package model

import "github.com/rymiyamoto/affiliate-api/model/common"

type (
	// Product struct
	Product struct {
		ID       int64
		ShopType common.ShopType
		Code     string
		Price    int64
	}

	// Products struct list
	Products []Product
)
