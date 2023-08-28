package main

import (
	"fmt"
)

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

type Describable interface {
	Description()
	DescriptionPrint()
}

func (i Item) String() string {
	if i.Discount > 0 {
		s := fmt.Sprintf("%s - %.2fTL (%.f tl indirimle %.2f tl)", i.Name, i.Price, i.Discount, calculatePrice(i))
		return s
	}
	s := fmt.Sprintf("%s - %.2fTL", i.Name, i.Price)
	return s
}

func (i Item) Format(f fmt.State, verb rune) {
	val := i.String()
	if verb == 81 {
		fmt.Fprint(f, val)
	}
}

func (i Item) DescriptionPrint() {
	i.Description()
}

func (i Item) Description() {
	if i.Discount > 0 {
		fmt.Printf("%Q\n", i)
		return
	}
	fmt.Printf("%Q\n", i)
}

func calculatePrice(item Item) float64 {
	if item.Discount > 0 {
		return item.Price - item.Discount
	}
	return item.Price
}

func totalPrice(items []Item) float64 {
	var total float64
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

func main() {
	items := [...]Item{
		{"Elma", 100, 10},
		{"Armut", 50, 0},
		{"Kitaplık", 200, 0},
		{"Koltuk", 150, 25},
	}

	items[0].DescriptionPrint()
	items[1].DescriptionPrint()
	items[2].DescriptionPrint()
	items[3].DescriptionPrint()
	totalPrice(items[:])
}
