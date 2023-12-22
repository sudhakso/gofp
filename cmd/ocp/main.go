package main

import "fmt"

type Color int

const (
	green Color = iota
	red
	blue
)

type Size int

const (
	small Size = iota
	big
)

// concrete type
type Product struct {
	name  string
	color Color
	size  Size
}

// concrete receiver type for filtering by features - anti-pattern as per OCP
// We have to change Filter type each time to extend the filtering capabilities.
type Filter struct {
}

func (f *Filter) filterByColor(p []Product, color Color) []*Product {
	fmt.Printf("Filtering by color (Old)\n")
	result := make([]*Product, 0)
	for i, v := range p {
		if v.color == color {
			result = append(result, &p[i])
		}
	}
	return result
}

func (f *Filter) filterBySize(p []Product, size Size) []*Product {
	fmt.Printf("Filtering by size (Old)\n")
	result := make([]*Product, 0)
	for i, v := range p {
		if v.size == size {
			result = append(result, &p[i])
		}
	}
	return result
}

// OCP compliant way

type FilterSpec interface {
	IsSatisfied(*Product) bool
}

type ColorFilter struct {
	color Color
}

func (cf *ColorFilter) IsSatisfied(prod *Product) bool {
	return cf.color == prod.color
}

type SizeFilter struct {
	size Size
}

func (sf *SizeFilter) IsSatisfied(product *Product) bool {
	return sf.size == product.size
}

func main() {
	fmt.Println("OCP Main...")

	apple := Product{"apple", green, small}
	tree := Product{"tree", green, big}
	house := Product{"house", red, big}

	products := []Product{apple, tree, house}

	// Calling filters - anti-pattern
	f := Filter{}
	gp := f.filterByColor(products, green)
	for _, v := range gp {
		fmt.Printf(" - %s is green\n", v.name)
	}

	sp := f.filterBySize(products, small)
	for _, v := range sp {
		fmt.Printf(" - %s is small\n", v.name)
	}

	// Doing the OCP compliant way
	cf := ColorFilter{green}
	fmt.Println("Filtering by Color (New)")
	for _, v := range products {
		if cf.IsSatisfied(&v) {
			fmt.Printf(" - %s is green\n", v.name)
		}
	}

	// Doing the OCP compliant way
	sf := SizeFilter{small}
	fmt.Println("Filtering by Size (New)")
	for _, v := range products {
		if sf.IsSatisfied(&v) {
			fmt.Printf(" - %s is small\n", v.name)
		}
	}
}
