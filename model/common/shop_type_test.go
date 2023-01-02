package common

import (
	"fmt"
	"testing"
)

func TestShopType_ToShopType(t *testing.T) {
	t.Parallel()

	exec := func(shop string, expect ShopType) {
		t.Run(fmt.Sprintf("when shop is %s", shop), func(t *testing.T) {
			t.Parallel()

			ret := ToShopType(shop)

			if ret != expect {
				t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", expect, ret))
			}
		})
	}

	tests := []struct {
		shop   string
		expect ShopType
	}{
		{shop: "unknown", expect: ShopTypeUnknown},
		{shop: "rakuten", expect: ShopTypeRakuten},
		{shop: "yahoo", expect: ShopTypeYahoo},
	}

	for _, test := range tests {
		exec(test.shop, test.expect)
	}
}
