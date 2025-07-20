package models

import "vending_machine/internal/enums"

type Item struct {
	ItemId int
	Type   enums.ProductType
	Price  int
}
