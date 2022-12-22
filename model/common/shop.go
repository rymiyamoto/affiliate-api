package common

type (
	ShopType int64
)

const (
	ShopTypeUnknown ShopType = iota
	ShopTypeRakuten
	ShopTypeYahoo
)

func ToShopType(shop string) ShopType {
	switch shop {
	case "rakuten":
		return ShopTypeRakuten
	case "yahoo":
		return ShopTypeYahoo
	default:
		return ShopTypeUnknown
	}
}
