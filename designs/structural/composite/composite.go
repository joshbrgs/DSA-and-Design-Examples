package composite

import (
	"fmt"
)

// Component Interface: This interface defines the common operations that both leaf and composite objects must implement. In our example, the Product interface defines the GetPrice method.

// Leaf: A leaf is a basic element that doesn't have any children. In our example, SingleProduct is a leaf that represents a single product.

// Composite: A composite is a container that can hold leaf elements or other composites. In our example, ProductBundle is a composite that can contain multiple products.

// Client: The client interacts with the components through the common interface. The client doesn't need to know whether it's dealing with a leaf or a composite. This allows the client to work with complex structures in a uniform way.

// Component interface
type Product interface {
	GetPrice() float64
}

// Leaf
type SingleProduct struct {
	price float64
}

func (p *SingleProduct) GetPrice() float64 {
	return p.price
}

// Composite
type ProductBundle struct {
	products []Product
}

func (b *ProductBundle) AddProduct(p Product) {
	b.products = append(b.products, p)
}

func (b *ProductBundle) GetPrice() float64 {
	total := 0.0
	for _, p := range b.products {
		total += p.GetPrice()
	}
	return total
}

func main() {
	// Create single products
	product1 := &SingleProduct{price: 25.0}
	product2 := &SingleProduct{price: 45.0}

	// Create a bundle and add single products
	bundle := &ProductBundle{}
	bundle.AddProduct(product1)
	bundle.AddProduct(product2)

	fmt.Printf("Total price of the bundle: $%.2f\n", bundle.GetPrice()) // Output: Total price of the bundle: $70.00
}
