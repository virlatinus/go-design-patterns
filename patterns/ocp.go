package patterns

import (
	"fmt"
	"github.com/pterm/pterm"
)

// OCP
// open for extension, close for modification
// specification and composite patterns

type OpenClose struct{}

//go:generate stringer -type=Size
type Size int

//go:generate stringer -type=Color
type Color int

const (
	Small Size = iota
	Medium
	Large
)

const (
	Red Color = iota
	Green
	Blue
)

type Product struct {
	Name string
	Color
	Size
}

type Products []Product

func (p Product) String() string {
	return fmt.Sprintf("%s (%s, %s)", p.Name, p.Color, p.Size)
}

func (ps Products) String() string {
	s := ""
	for _, p := range ps {
		s += fmt.Sprintln(p)
	}
	return s
}

type Specification interface {
	IsSatisfied(p Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

func (c ColorSpecification) IsSatisfied(p Product) bool {
	return c.color == p.Color
}

func (s SizeSpecification) IsSatisfied(p Product) bool {
	return s.size == p.Size
}

// AndSpecification composite specification
type AndSpecification struct {
	first, second Specification
}

func (c AndSpecification) IsSatisfied(p Product) bool {
	return c.first.IsSatisfied(p) && c.second.IsSatisfied(p)
}

type Filter struct{}

func (f *Filter) Filter(products []Product, spec Specification) Products {
	result := make(Products, 0)
	for i, v := range products {
		if spec.IsSatisfied(v) {
			result = append(result, products[i])
		}
	}
	return result
}

func (o OpenClose) Run() {
	pterm.DefaultBasicText.Println()
	pterm.DefaultBox.Println("SOLID Open Close Principle")

	apple := Product{"Apple", Green, Small}
	tree := Product{"Tree", Green, Large}
	house := Product{"House", Blue, Large}

	products := Products{apple, tree, house}

	fmt.Println("Initial data:")
	fmt.Println(products)

	pterm.DefaultBasicText.Println("Filtering by color (" + pterm.Green("green") + "):")
	greenSpec := ColorSpecification{Green}
	filter := Filter{}
	greenProducts := filter.Filter(products, greenSpec)
	fmt.Println(greenProducts)

	fmt.Println("Filtering by size (large):")
	largeSpec := SizeSpecification{Large}
	largeProducts := filter.Filter(products, largeSpec)
	fmt.Println(largeProducts)

	fmt.Println("Filtering by size and color (large, green):")
	lgSpec := AndSpecification{largeSpec, greenSpec}
	largeGreenProducts := filter.Filter(products, lgSpec)
	fmt.Println(largeGreenProducts)
}
