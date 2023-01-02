package common

import (
	"fmt"
	"testing"
)

func TestShopType_ToShopType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		shop   string
		expect ShopType
	}{
		{shop: "unknown", expect: ShopTypeUnknown},
		{shop: "rakuten", expect: ShopTypeRakuten},
		{shop: "yahoo", expect: ShopTypeYahoo},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("when shop is %s", test.shop), func(t *testing.T) {
			ret := ToShopType(test.shop)

			if ret != test.expect {
				t.Error(fmt.Errorf("expected same value. expect: %+v, actual: %+v", test.expect, ret))
			}
		})
	}
}
