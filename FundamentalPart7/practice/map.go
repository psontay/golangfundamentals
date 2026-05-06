package practice

import (
	"errors"
	"fmt"
)

var ErrNotEnoughStock = errors.New("not enough stock")
var ErrNoSuchProduct = errors.New("no such product")

type Product struct {
	Price    float64
	Quantity int
}

type mapSProduct map[string]*Product

func UpdateStock(inventory mapSProduct, name string, change int) error {
	product, ok := inventory[name]
	if !ok {
		return ErrNoSuchProduct
	}
	newQuantity := product.Quantity + change
	if newQuantity < 0 {
		return ErrNotEnoughStock
	}
	product.Quantity = newQuantity
	fmt.Printf("Update successfully: %s, New Stock: %d\n", name, product.Quantity)
	return nil
}
