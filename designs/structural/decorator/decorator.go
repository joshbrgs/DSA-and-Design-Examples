package decorator

import (
	"fmt"
)

// Let's break down the key components of the Decorator Pattern:

// Component: This is the base interface or class that defines the object to which additional functionalities can be added. In the example above, Coffee is the component interface.

// ConcreteComponent: This is the base class that implements the Component interface. In the example, SimpleCoffee is the concrete component.

// Decorator: This is the abstract class that wraps around the Component interface and provides additional functionalities. In the example, CoffeeDecorator is the decorator class.

// ConcreteDecorator: These are the concrete classes that extend the Decorator class and add specific functionalities. In the example, MilkDecorator and SugarDecorator are concrete decorators.

// Client: This is the class that uses the Component interface and can add functionalities using decorators. In the example, main is the client code.

// Component interface
type Coffee interface {
	Cost() float64
	Description() string
}

// ConcreteComponent
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 5.0
}

func (c *SimpleCoffee) Description() string {
	return "Simple coffee"
}

// Decorator
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}

// ConcreteDecorators
type MilkDecorator struct {
	*CoffeeDecorator
}

func NewMilkDecorator(c Coffee) *MilkDecorator {
	return &MilkDecorator{&CoffeeDecorator{c}}
}

func (d *MilkDecorator) Cost() float64 {
	return d.CoffeeDecorator.Cost() + 1.0
}

func (d *MilkDecorator) Description() string {
	return d.CoffeeDecorator.Description() + ", milk"
}

type SugarDecorator struct {
	*CoffeeDecorator
}

func NewSugarDecorator(c Coffee) *SugarDecorator {
	return &SugarDecorator{&CoffeeDecorator{c}}
}

func (d *SugarDecorator) Cost() float64 {
	return d.CoffeeDecorator.Cost() + 0.5
}

func (d *SugarDecorator) Description() string {
	return d.CoffeeDecorator.Description() + ", sugar"
}

func main() {
	var coffee Coffee = &SimpleCoffee{}
	fmt.Println(coffee.Description(), ":", coffee.Cost()) // Simple coffee : 5

	coffee = NewMilkDecorator(coffee)
	fmt.Println(coffee.Description(), ":", coffee.Cost()) // Simple coffee, milk : 6

	coffee = NewSugarDecorator(coffee)
	fmt.Println(coffee.Description(), ":", coffee.Cost()) // Simple coffee, milk, sugar : 6.5
}
