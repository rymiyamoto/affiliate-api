package repository

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rymiyamoto/affiliate-api/model"
	"github.com/rymiyamoto/affiliate-api/model/common"
)

func TestProduct_NewProduct(t *testing.T) {
	t.Parallel()
	t.Run("NewProduct", func(t *testing.T) {
		expect := &Product{}

		ret := NewProduct()

		if diff := cmp.Diff(ret, expect); diff != "" {
			t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", expect, ret))
		}
	})
}

func TestProduct_ByShopType(t *testing.T) {
	t.Parallel()

	t.Run("success case", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			shop   common.ShopType
			expect *model.Products
		}{
			{
				shop: common.ShopTypeRakuten,
				expect: &model.Products{
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
				},
			},
			{
				shop: common.ShopTypeYahoo,
				expect: &model.Products{
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
				},
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("when shop_type is %d", test.shop), func(t *testing.T) {
				t.Parallel()

				ret, err := NewProduct().ByShopType(test.shop)

				if err != nil {
					t.Error(fmt.Errorf("expected didn't has error. err: %w", err))
				}

				if diff := cmp.Diff(ret, test.expect); diff != "" {
					t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", test.expect, ret))
				}
			})
		}
	})

	t.Run("fail case", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			shop   common.ShopType
			expect error
		}{
			{
				shop:   common.ShopTypeUnknown,
				expect: fmt.Errorf("not define type. shop_type: %d", common.ShopTypeUnknown),
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("when shop_type is %d", test.shop), func(t *testing.T) {
				t.Parallel()

				ret, err := NewProduct().ByShopType(test.shop)

				if ret != nil {
					t.Error(fmt.Errorf("expected didn't has ret. ret: %+v", ret))
				}

				if diff := cmp.Diff(err.Error(), test.expect.Error()); diff != "" {
					t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", test.expect, err))
				}
			})
		}
	})
}
