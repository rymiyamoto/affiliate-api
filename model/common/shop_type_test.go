package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, ToShopType(test.shop), test.expect)
	}
}
