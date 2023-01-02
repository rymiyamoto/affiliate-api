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

		exec := func(shop common.ShopType, expect *model.Products) {
			t.Run(fmt.Sprintf("when shop_type is %d", shop), func(t *testing.T) {
				t.Parallel()

				ret, err := NewProduct().ByShopType(shop)

				if err != nil {
					t.Error(fmt.Errorf("expected didn't has error. err: %w", err))
				}

				if diff := cmp.Diff(ret, expect); diff != "" {
					t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", expect, ret))
				}
			})
		}

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
						Code:     "8e27892c206a7118a74445c5a0825a41",
						Price:    8000,
					},
					{
						ID:       11,
						ShopType: common.ShopTypeRakuten,
						Code:     "a3da4c917cdfdbce5577f6923533f003",
						Price:    5000,
					},
					{
						ID:       12,
						ShopType: common.ShopTypeRakuten,
						Code:     "1c3d37510cfe620aa200862886a6502c",
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
						Code:     "4949776441067",
						Price:    7000,
					},
					{
						ID:       21,
						ShopType: common.ShopTypeYahoo,
						Code:     "4949776342067",
						Price:    4000,
					},
					{
						ID:       22,
						ShopType: common.ShopTypeYahoo,
						Code:     "4948872325134",
						Price:    500,
					},
				},
			},
		}

		for _, test := range tests {
			exec(test.shop, test.expect)
		}
	})

	t.Run("fail case", func(t *testing.T) {
		t.Parallel()

		exec := func(shop common.ShopType, expect error) {
			t.Run(fmt.Sprintf("when shop_type is %d", shop), func(t *testing.T) {
				t.Parallel()

				ret, err := NewProduct().ByShopType(shop)

				if ret != nil {
					t.Error(fmt.Errorf("expected didn't has ret. ret: %+v", ret))
				}

				if diff := cmp.Diff(err.Error(), expect.Error()); diff != "" {
					t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", expect, err))
				}
			})
		}

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
			exec(test.shop, test.expect)
		}
	})
}
